package gredis

import (
	"bytes"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/tmpl"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

// mkDirAll clear all files in rootDir and create the dir tree:
//
//	rootDir/
//	├── dbm/
//	├── dbnames/
//	├── cache/
//	└── internal/
//	|   ├── genmysql/
//	|   └── genredis/
//	└── params/
func mkDirAll(rootDir string) error {
	if err := os.RemoveAll(rootDir); err != nil && os.IsNotExist(err) {
		return err
	}
	//if err := os.MkdirAll(path.Join(rootDir, "dbm"), 0755); err != nil && os.IsExist(err) {
	//	return err
	//}
	if err := os.MkdirAll(path.Join(rootDir, "dbnames"), 0755); err != nil && os.IsExist(err) {
		return err
	}
	if err := os.MkdirAll(path.Join(rootDir, "cache"), 0755); err != nil && os.IsExist(err) {
		return err
	}
	if err := os.MkdirAll(path.Join(rootDir, "internal", "genmysql"), 0755); err != nil && os.IsExist(err) {
		return err
	}
	if err := os.MkdirAll(path.Join(rootDir, "internal", "genredis"), 0755); err != nil && os.IsExist(err) {
		return err
	}
	if err := os.MkdirAll(path.Join(rootDir, "internal", "gengocache"), 0755); err != nil && os.IsExist(err) {
		return err
	}
	if err := os.MkdirAll(path.Join(rootDir, "params"), 0755); err != nil && os.IsExist(err) {
		return err
	}
	return nil
}

func genFileTmplName(tmplName, outputFile string, params interface{}) {
	tpl, err := tmpl.Load(tmplName)
	if err != nil {
		log.Println("[Error] Loading template:", tmplName)
		panic(err)
	}
	if err = genFileTmpl(tpl, outputFile, params, nil, nil); err != nil {
		log.Println("[Error] genFileTmpl:", err.Error())
	}
}

func genFileTmpl(tmpl *template.Template, outputFile string, params interface{}, head, tail []byte) (err error) {
	buf := bytes.NewBuffer(make([]byte, 0, 1024*1024))

	if err = genWriterTmpl(tmpl, buf, params, head, tail); err != nil {
		return
	}
	if err = os.Remove(outputFile); err != nil && !os.IsNotExist(err) {
		log.Println("[Warn] remove", outputFile, ">>>", err.Error())
	}
	genCode, err := beforeWriteFile(buf.Bytes())
	if err != nil {
		log.Println("[Error] beforeWriteFile:", outputFile, err.Error(), ">>>\n", buf.String())
		return
	}
	if err = os.WriteFile(outputFile, genCode, 0644); err != nil {
		return
	}
	return
}

func genWriterTmpl(tmpl *template.Template, w io.Writer, params interface{}, head, tail []byte) (err error) {
	if len(head) > 0 {
		if _, err = w.Write(head); err != nil {
			log.Println("[Error] write file head:", err.Error(), ">>>\n", string(head))
			return
		}
	}
	if err = tmpl.Execute(w, params); err != nil {
		log.Println("[Error] Executing template:", tmpl.Name())
		return
	}
	if len(tail) > 0 {
		if _, err = w.Write(tail); err != nil {
			log.Println("[Error] write file tail:", err.Error(), ">>>\n", string(tail))
			return
		}
	}
	return
}

func beforeWriteFile(rawBuf []byte) (buf []byte, err error) {
	// 特殊情况： 清理导入了但是没有用到的fmt包
	if bytes.Contains(rawBuf, []byte(`"fmt"`)) && !bytes.Contains(rawBuf, []byte(`fmt.`)) {
		rawBuf = bytes.ReplaceAll(rawBuf, []byte(`"fmt"`), []byte{})
	}
	return format.Source(rawBuf)
}

func copyGoFile(srcFile, outputFile string) (err error) {
	body, err := tmpl.LoadFile(srcFile)
	if err != nil {
		return
	}
	if err = ioutil.WriteFile(outputFile, body, 0644); err != nil {
		return
	}
	return
}
