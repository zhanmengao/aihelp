package ghash

import (
	"fmt"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/gredis/internal/gredistyp"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/tmpl"
	"path"
	"strings"
	"text/template"
)

type genRedisHash struct {
	keyFmt     string
	keyParam   []*gredistyp.TKeyParam
	fieldFmt   string
	fieldParam []*gredistyp.TKeyParam
	pbName     string
	dbName     string
	genMGet    bool
	buildTag   string
}

// NewRedisHash WorldData:%s:%s,Platform,Channel|UserData:%s,UID
func NewRedisHash(dbName, pbName, buildTag string, paramList []string, fm map[string]string) *genRedisHash {
	ret := &genRedisHash{
		dbName:   dbName,
		pbName:   pbName,
		buildTag: buildTag,
	}
	//WorldData:%s:%s,Platform,Channel
	keyParam := strings.Split(paramList[0], ",")
	ret.keyFmt = keyParam[0]
	ret.keyParam = make([]*gredistyp.TKeyParam, 0, len(keyParam)-1)
	for _, param := range keyParam[1:] {
		//获取参数类型
		param := &gredistyp.TKeyParam{
			ParamName: param,
			ParamType: fm[param],
		}
		ret.keyParam = append(ret.keyParam, param)
	}

	//UserData:%s,UID
	fieldParam := strings.Split(paramList[1], ",")
	ret.fieldFmt = fieldParam[0]
	ret.fieldParam = make([]*gredistyp.TKeyParam, 0, len(fieldParam)-1)
	for _, param := range fieldParam[1:] {
		//获取参数类型
		param := &gredistyp.TKeyParam{
			ParamName: param,
			ParamType: fm[param],
		}
		ret.fieldParam = append(ret.fieldParam, param)
	}
	ret.genMGet = len(keyParam) > 0 && paramList[len(paramList)-1] == gredistyp.MultiMark
	return ret
}

func (s *genRedisHash) GetTemplate() (tpl *template.Template, err error) {
	if tpl, err = tmpl.Load("db/redis_hash.tmpl"); err != nil {
		return
	}
	return
}

func (s *genRedisHash) GetGenInfo() *gredistyp.TGenRedisInfo {
	params := make([]*gredistyp.TKeyParam, 0, len(s.keyParam)+len(s.fieldParam))
	params = append(params, s.keyParam...)
	params = append(params, s.fieldParam...)
	ret := &gredistyp.TGenRedisInfo{
		DBName:        s.dbName,
		Type:          "hash",
		StructName:    s.pbName,
		Params:        params,
		KeyFmt:        s.keyFmt,
		KeyParams:     s.keyParam,
		FieldFmt:      s.fieldFmt,
		FieldParams:   s.fieldParam,
		MultiMark:     s.genMGet,
		MParams:       s.genMGet && len(s.keyParam)+len(s.fieldParam) > 1,
		BuildTagLine1: "",
		BuildTagLine2: "",
	}
	if s.buildTag != "" {
		ret.BuildTagLine1 = fmt.Sprintf("//go:build %s", s.buildTag)
		ret.BuildTagLine2 = fmt.Sprintf("build %s", s.buildTag)
	} else {
		ret.BuildTagLine1 = "//"
		ret.BuildTagLine2 = "//"
	}
	return ret
}

func (s *genRedisHash) GetOutput(pbgoDir, dbRoot string) string {
	genFileName := fmt.Sprintf("%s_%s.db.go", strings.ToLower(s.dbName), strings.ToLower(s.pbName))
	filePath := path.Join(dbRoot, "internal", "genredis", genFileName)
	return filePath
}
