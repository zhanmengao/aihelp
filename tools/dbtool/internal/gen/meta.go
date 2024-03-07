package gen

import (
	"bytes"
	"fmt"
	"forevernine.com/midplat/base_server/tools/f9/dbtool/internal/dbtype/tmpl"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"forevernine.com/midplat/base_server/tools/f9/dbtool/internal/dbtype/gmysql"
	"forevernine.com/midplat/base_server/tools/f9/dbtool/internal/dbtype/gredis"
)

type DBMeta struct {
	tmpTypeMap     map[string]*ast.TypeSpec
	tmpMsgComments []string
	tmpDBStart     bool

	PkgPath   string
	PbPkg     string
	GenPkg    string
	outputDir string
	comments  map[string]string
	// StructFields[structName][fieldName] = fieldType
	StructFields map[string]map[string]string
	buildTag     string
}

func NewDBMeta() *DBMeta {
	return &DBMeta{
		tmpMsgComments: make([]string, 0, 5),
		tmpTypeMap:     make(map[string]*ast.TypeSpec, 5),
		comments:       make(map[string]string, 5),
		StructFields:   make(map[string]map[string]string, 5),
	}
}

// ParseComments 解析注释获取db结构声明
func (d *DBMeta) ParseComments(pbGoPath string) (err error) {
	fSet := token.NewFileSet()
	//把所有文件在ast里跑一遍
	pkg, err := parser.ParseDir(fSet, pbGoPath, func(info os.FileInfo) bool {
		return true
	}, parser.ParseComments)
	if err != nil {
		return
	}
	for k, v := range pkg {
		d.PbPkg = k
		ast.Walk(d, v)
	}
	return
}

func (d *DBMeta) Generate(pbgoDir string) {
	commentsKeys := make([]string, 0, len(d.comments))
	for k := range d.comments {
		commentsKeys = append(commentsKeys, k)
	}
	sort.Strings(commentsKeys)
	var genRedis = gredis.NewRedisGenerate(d.PkgPath, d.GenPkg, d.outputDir, pbgoDir, d.PbPkg)
	var genMysql = gmysql.NewMysqlGenerate(d.PbPkg, pbgoDir)
	//先写个inject头部
	injectBuffer := bytes.NewBuffer(make([]byte, 0, 4096))
	injectBuffer.Write(d.LoadInjectHead())

	for _, pbName := range commentsKeys {
		desc := d.comments[pbName]
		ss := strings.Split(desc, "|")
		paramList := make([]string, 0, 5)
		var dbType string
		var dbName string
		for i, v := range ss {
			if i == 0 {
				//@db:hash
				dbType = strings.Split(v, ":")[1]
			} else if i == 1 {
				dbName = v
			} else {
				paramList = append(paramList, v)
			}
		}
		//@db:hash
		switch dbType {
		case "hash", "string":
			genRedis.Generate(dbName, dbType, pbName, d.buildTag, desc, paramList, d.StructFields[pbName], injectBuffer)
		case "mysql":
			genMysql.Generate(dbName, dbType, pbName, desc, paramList, d.StructFields[pbName])
		}
	}
	genRedis.End()
	genMysql.End()

	injectGo := filepath.Join(pbgoDir, "db.inject.go")
	if err := ioutil.WriteFile(injectGo, injectBuffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func (d *DBMeta) LoadInjectHead() []byte {
	redisOutHead, err := tmpl.Load("inject/inject_head.tmpl")
	if err != nil {
		panic(err)
	}
	BuildTagLine1 := "//"
	BuildTagLine2 := "//"
	if d.buildTag != "" {
		BuildTagLine1 = fmt.Sprintf("//go:build %s", d.buildTag)
		BuildTagLine2 = fmt.Sprintf("build %s", d.buildTag)
	}
	headBuf := bytes.NewBuffer(make([]byte, 0, 4096))
	if err = redisOutHead.Execute(headBuf, map[string]interface{}{
		"PbPkg":         d.PbPkg,
		"BuildTagLine1": BuildTagLine1,
		"BuildTagLine2": BuildTagLine2,
	}); err != nil {
		panic(err)
	}
	return headBuf.Bytes()
}
