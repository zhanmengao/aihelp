package gstring

import (
	"fmt"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/gredis/internal/gredistyp"
	"github.com/zhanmengao/aihelp/tools/dbtool/internal/dbtype/tmpl"
	"path"
	"strings"
	"text/template"
)

type genRedisString struct {
	keyFmt   string
	keyParam []*gredistyp.TKeyParam
	dbName   string
	pbName   string
	genMGet  bool
	buildTag string
}

// NewRedisString UserData:%s:%s,Platform,UID
func NewRedisString(dbName, pbName, buildTag string, paramList []string, fm map[string]string) *genRedisString {
	ret := &genRedisString{
		dbName:   dbName,
		pbName:   pbName,
		buildTag: buildTag,
	}
	keyParam := strings.Split(paramList[0], ",")
	ret.keyParam = make([]*gredistyp.TKeyParam, 0, len(keyParam)-1)
	ret.keyFmt = keyParam[0]
	for _, param := range keyParam[1:] {
		//获取参数类型
		param := &gredistyp.TKeyParam{
			ParamName: param,
			ParamType: fm[param],
		}
		ret.keyParam = append(ret.keyParam, param)
	}
	ret.genMGet = len(keyParam) > 0 && paramList[len(paramList)-1] == gredistyp.MultiMark
	return ret
}

func (s *genRedisString) GetTemplate() (tpl *template.Template, err error) {
	return tmpl.Load("db/redis_string.tmpl")
}

func (s *genRedisString) GetGenInfo() *gredistyp.TGenRedisInfo {
	ret := &gredistyp.TGenRedisInfo{
		DBName:        s.dbName,
		Type:          "string",
		StructName:    s.pbName,
		Params:        s.keyParam,
		KeyFmt:        s.keyFmt,
		KeyParams:     s.keyParam,
		MultiMark:     s.genMGet,
		MParams:       s.genMGet && len(s.keyParam) > 1,
		BuildTagLine1: "",
		BuildTagLine2: "",
		FieldFmt:      "",
		FieldParams:   nil,
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

func (s *genRedisString) GetOutput(pbgoDir, dbRoot string) string {
	genFileName := fmt.Sprintf("%s_%s.db.go", strings.ToLower(s.dbName), strings.ToLower(s.pbName))
	filePath := path.Join(dbRoot, "internal", "genredis", genFileName)
	return filePath
}
