package gmysql

import (
	"bytes"
	"log"
	"os"
)

type genMysql struct {
	pbGoDir   string
	pbPackage string
	pf        *os.File
	buf       *bytes.Buffer
	info      *genMysqlInfo
}

// NewMysqlGenerate @db:mysql|FLS|t_chat_msg
func NewMysqlGenerate(pbPackage, pbGoDir string) *genMysql {
	ret := &genMysql{
		pbGoDir:   pbGoDir,
		pbPackage: pbPackage,
		buf:       bytes.NewBuffer(make([]byte, 0, 1024*1024)),
		info: &genMysqlInfo{
			PbPackage: pbPackage,
		},
	}
	//pf, err := os.OpenFile(path.Join(pbGoDir, "gen_mysql.db.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//ret.pf = pf
	//headTpl, err := tmpl.Load("mysql_inject_head.tmpl")
	//if err != nil {
	//	panic(err)
	//}
	//err = headTpl.Execute(ret.buf, ret.info)
	//if err != nil {
	//	panic(err)
	//}
	return ret
}

func (r *genMysql) Generate(dbName, dbType, pbName, desc string, paramList []string, metaInfo map[string]string) {
	//tableName := paramList[0]
	//bodyTpl, err := tmpl.Load("mysql_inject_body.tmpl")
	//if err != nil {
	//	panic(err)
	//}
	//r.info.Comment = desc
	//r.info.StructName = pbName
	//r.info.TableName = tableName
	//if err = bodyTpl.Execute(r.buf, r.info); err != nil {
	//	r.writeFile(r.buf.Bytes())
	//	panic(err)
	//}
}

func (r *genMysql) End() {
	//var genCode []byte
	//defer r.pf.Close()
	//defer func() {
	//	r.writeFile(genCode)
	//}()
	//genCode, err := format.Source(r.buf.Bytes())
	//if err != nil {
	//	log.Println("[Error] gofmt file:", r.buf.Bytes(), err.Error(), ">>>\n", r.buf.String())
	//	return
	//}
}

func (r *genMysql) writeFile(genCode []byte) (err error) {
	if _, err = r.pf.Write(genCode); err != nil {
		log.Println(string(genCode))
	}
	return
}
