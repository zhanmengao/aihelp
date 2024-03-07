// Package tmpl 内嵌了db生成工具所需的模板文件，并提供了响应的文件加载方法
package tmpl

import (
	"embed"
	"path"
	"path/filepath"
	"text/template"
)

//go:embed template/db/*.tmpl template/dbm/*.tmpl template/cache/* template/*.tmpl template/inject/*
var fs embed.FS

// Load 通过文件名加载模板文件
func Load(fn string) (tpl *template.Template, err error) {
	data, err := fs.ReadFile(path.Join("template", fn))
	if err != nil {
		return
	}
	tplName := filepath.Base(fn)
	ext := filepath.Ext(tplName)
	tplName = tplName[:len(tplName)-len(ext)-1]
	tpl, err = template.New(tplName).Parse(string(data))
	if err != nil {
		return
	}
	return
}

func LoadFile(fn string) (body []byte, err error) {
	body, err = fs.ReadFile(path.Join("template", fn))
	if err != nil {
		return
	}
	return
}
