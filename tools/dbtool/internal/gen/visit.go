package gen

import (
	"fmt"
	"go/ast"
	"strings"
)

func (d *DBMeta) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.Package:
		d.PbPkg = n.Name
		return d
	case *ast.File:
		return d
	case *ast.GenDecl:
		if n.Doc == nil {
			// 没有类型注释，跳过
			return d
		}
		d.findMarkedType(n)
		return d
	case *ast.TypeSpec:
		if d.tmpDBStart {
			d.comments[n.Name.Name] = strings.Join(d.tmpMsgComments, "")
			d.tmpDBStart = false
			d.tmpMsgComments = make([]string, 0)
		}
		return d
	}

	return d
}

// findMarkedType 找到被 @db 标记的结构体类型声明
// 并记录对应的标记内容和该结构体的字段信息
func (d *DBMeta) findMarkedType(n *ast.GenDecl) {
	for _, comment := range n.Doc.List {
		if !strings.Contains(comment.Text, "@db:") {
			// 不是生成标记，跳过
			continue
		}
		d.tmpDBStart = true
		if d.tmpMsgComments == nil {
			d.tmpMsgComments = make([]string, 0, 1)
		}
		d.tmpMsgComments = append(d.tmpMsgComments, comment.Text)
		//抓取类型信息
		if len(n.Specs) != 1 {
			// 不是仅包含一个类型定义，跳过
			continue
		}
		ss := n.Specs[0]
		switch tn := ss.(type) {
		case *ast.TypeSpec:
			structName := tn.Name.Name // 被标记的结构体名称
			st, ok := tn.Type.(*ast.StructType)
			if !ok {
				// 不是结构体声明，跳过
				continue
			}
			d.tmpTypeMap[structName] = tn
			d.StructFields[structName] = make(map[string]string, len(st.Fields.List))
			for _, v := range st.Fields.List {
				fieldName := v.Names[0].Name
				if ident, ok := v.Type.(*ast.Ident); ok {
					fieldType := ident.Name
					switch ident.Name {
					case "string", "int32", "int64", "uint32", "uint64", "float32", "float64", "float":
					default:
						fieldType = fmt.Sprintf("%s.%s", d.PbPkg, fieldType)
					}
					d.StructFields[structName][fieldName] = fieldType
				}
			}
		}
	}
}
