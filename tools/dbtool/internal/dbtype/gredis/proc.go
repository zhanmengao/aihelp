package gredis

import (
	"bytes"
	"fmt"
	"forevernine.com/midplat/base_libs/errors"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/gredis/internal/ghash"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/gredis/internal/gredistyp"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/gredis/internal/gstring"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/tmpl"
	"io"
	"log"
	"path"
	"strings"
	"text/template"
)

type RedisGenerate struct {
	pbPkgPath  string
	genPkgPath string
	outputDir  string
	pbGoDir    string
	pbPkg      string
	dbNames    []*uniqVal
	redisDBs   []*uniqVal
	mParams    bool
}

func NewRedisGenerate(pbPkgPath, genPkgPath, outputDir, pbGoDir, pbPkg string) *RedisGenerate {
	gen := &RedisGenerate{
		pbPkgPath:  pbPkgPath,
		genPkgPath: genPkgPath,
		outputDir:  outputDir,
		pbGoDir:    pbGoDir,
		pbPkg:      pbPkg,
	}

	if err := mkDirAll(outputDir); err != nil {
		log.Println("mkDirAll:", err.Error())
		panic(err)
	}
	return gen
}

func (r *RedisGenerate) Generate(dbName, dbType, pbName, buildTag, desc string, paramList []string, metaInfo map[string]string, injectW io.Writer) {
	//redis
	var err error
	g := r.getGenerator(dbName, dbType, pbName, buildTag, paramList, metaInfo)
	if g == nil {
		return
	}
	var tpl *template.Template
	if tpl, err = g.GetTemplate(); err != nil {
		panic(err)
	}
	genInfo := g.GetGenInfo()
	genInfo.SrcComment = desc
	r.dbNames = uniqAdd(r.dbNames, genInfo.DBName, nil, nil)
	r.redisDBs = uniqAdd(r.redisDBs, genInfo.DBName, nil, genInfo)
	r.mParams = r.mParams || genInfo.MParams
	head := r.loadRedisHead(genInfo.MParams, genInfo.BuildTagLine1, genInfo.BuildTagLine2)
	if err = genFileTmpl(tpl, g.GetOutput(r.pbGoDir, r.outputDir), genInfo, head, nil); err != nil {
		log.Println("[Error] genFileTmpl:", err.Error())
	}
	//cache
	cacheTmpl, err := tmpl.Load("cache/cache_body.tmpl")
	if err != nil {
		panic(err)
	}
	of := path.Join(r.outputDir, "internal", "gengocache",
		fmt.Sprintf("%s_%s.db.go", strings.ToLower(dbName), strings.ToLower(pbName)))
	if err = genFileTmpl(cacheTmpl, of, genInfo, r.loadCacheHead(genInfo.MParams, genInfo.BuildTagLine1, genInfo.BuildTagLine2), nil); err != nil {
		log.Println("[Error] genFileTmpl:", err.Error())
	} else {
		log.Println("gen ", of, "success")
	}
	//inject到pb包里
	injectTmpl, err := tmpl.Load("inject/inject_body.tmpl")
	if err != nil {
		panic(err)
	}
	if err = genWriterTmpl(injectTmpl, injectW, genInfo, nil, nil); err != nil {
		log.Println("[Error] genFileTmpl:", err.Error())
	} else {
		log.Println("gen ", of, "success")
	}
}

func (r *RedisGenerate) End() {
	if len(r.redisDBs) > 0 {
		genFileTmplName("db/redis_define.tmpl", path.Join(r.outputDir, "internal", "genredis", "define.db.go"), r.redisDBs)
		genFileTmplName("cache/redis_define.tmpl", path.Join(r.outputDir, "internal", "gengocache", "define.db.go"), r.redisDBs)
		genFileTmplName("redis_params.tmpl", path.Join(r.outputDir, "params", "redis_params.db.go"), r.redisDBs)
		fullParams := map[string]interface{}{
			"PbPkgPath":  r.pbPkgPath,
			"GenPkgPath": r.genPkgPath,
			"DBs":        r.redisDBs,
			"MParams":    r.mParams,
		}
		genFileTmplName("redis_factory.tmpl", path.Join(r.outputDir, "factory.db.go"), fullParams)

		// dbm
		//genFileTmplName("dbm/redis_api.tmpl", path.Join(r.outputDir, "dbm", "redis_api.dbm.go"), nil)
		//genFileTmplName("dbm/redis_meta.tmpl", path.Join(r.outputDir, "dbm", "redis_meta.dbm.go"), fullParams)
		//genFileTmplName("dbm/redis_parse_util.tmpl", path.Join(r.outputDir, "dbm", "redis_parse_util.dbm.go"), nil)
		//genFileTmplName("dbm/redis_fill_util.tmpl", path.Join(r.outputDir, "dbm", "redis_fill_util.dbm.go"), nil)
		//genFileTmplName("dbm/redis_fill_util_low.tmpl", path.Join(r.outputDir, "dbm", "redis_fill_util_low.dbm.go"), nil)
		//for _, eachDB := range r.redisDBs {
		//	fileName := fmt.Sprintf("redis_%s_wrapper.dbm.go", strings.ToLower(eachDB.Key))
		//	genFileTmplName("dbm/redis_wrapper.tmpl", path.Join(r.outputDir, "dbm", fileName), map[string]interface{}{
		//		"PbPkgPath":  r.pbPkgPath,
		//		"GenPkgPath": r.genPkgPath,
		//		"Slice":      eachDB.Slice,
		//	})
		//}
	}
	genFileTmplName("dbnames.tmpl", path.Join(r.outputDir, "dbnames", "dbnames.db.go"), r.dbNames)
	genFileTmplName("cache.tmpl", path.Join(r.outputDir, "cache", "cache.db.go"), nil)
	//直接拷贝go文件
	if err := copyGoFile("cache/vars.go", path.Join(r.outputDir, "internal", "gengocache", "vars.go")); err != nil {
		panic(err)
	}
}

// GetGenerator @db:hash|FLS|WorldData:%s:%s,Platform,Channel|UserData:%s,UID
func (r *RedisGenerate) getGenerator(dbName, dbType, pbName, buildTag string, paramList []string, metaInfo map[string]string) gredistyp.IDBGenRedisType {
	//@db:hash
	switch dbType {
	case "string":
		return gstring.NewRedisString(dbName, pbName, buildTag, paramList, metaInfo)
	case "hash":
		return ghash.NewRedisHash(dbName, pbName, buildTag, paramList, metaInfo)
	default:
		panic(errors.NewError(0, "").Format("%s : %s not allow ", pbName, dbType))
	}
}

func (r *RedisGenerate) loadRedisHead(useMParams bool, tagLine1, tagLine2 string) []byte {
	redisOutHead, err := tmpl.Load("db/redis_head.tmpl")
	if err != nil {
		panic(err)
	}
	headBuf := bytes.NewBuffer(make([]byte, 0, 4096))
	if err = redisOutHead.Execute(headBuf, map[string]interface{}{
		"PbPkgPath":     r.pbPkgPath,
		"GenPkgPath":    r.genPkgPath,
		"MParams":       useMParams,
		"BuildTagLine1": tagLine1,
		"BuildTagLine2": tagLine2,
	}); err != nil {
		panic(err)
	}
	return headBuf.Bytes()
}

func (r *RedisGenerate) loadCacheHead(useMParams bool, tagLine1, tagLine2 string) []byte {
	redisOutHead, err := tmpl.Load("cache/cache_head.tmpl")
	if err != nil {
		panic(err)
	}
	headBuf := bytes.NewBuffer(make([]byte, 0, 4096))
	if err = redisOutHead.Execute(headBuf, map[string]interface{}{
		"PbPkgPath":     r.pbPkgPath,
		"GenPkgPath":    r.genPkgPath,
		"MParams":       useMParams,
		"BuildTagLine1": tagLine1,
		"BuildTagLine2": tagLine2,
	}); err != nil {
		panic(err)
	}
	return headBuf.Bytes()
}
