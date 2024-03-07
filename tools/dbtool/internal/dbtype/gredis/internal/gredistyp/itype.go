package gredistyp

import "text/template"

const (
	DBTypeRedis = "redis"
	DBTypeMysql = "mysql"
)

const (
	MultiMark = "#multi"
)

type TKeyParam struct {
	ParamType string
	ParamName string
}

type TGenRedisInfo struct {
	DBName        string
	Type          string
	SrcComment    string
	StructName    string
	Params        []*TKeyParam
	KeyFmt        string
	KeyParams     []*TKeyParam
	FieldFmt      string
	FieldParams   []*TKeyParam
	MultiMark     bool
	MParams       bool
	BuildTagLine1 string
	BuildTagLine2 string
}

type IDBGenRedisType interface {
	GetTemplate() (tpl *template.Template, err error)
	GetGenInfo() *TGenRedisInfo
	GetOutput(pbgoDir, dbRoot string) string
}
