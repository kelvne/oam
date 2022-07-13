package astutils

import (
	"encoding/json"
	"go/ast"
)

func mapIfJSON(line string) (m map[string]interface{}, isJSON bool) {
	if err := json.Unmarshal([]byte(line), &m); err == nil {
		isJSON = true
	}
	return
}

func getEndIdent(expr ast.Expr) *ast.Ident {
	if star, ok := expr.(*ast.StarExpr); ok {
		if ident, ok := star.X.(*ast.Ident); ok {
			return ident
		}
		return getEndIdent(star.X)
	}
	if sel, ok := expr.(*ast.SelectorExpr); ok {
		if ident, ok := sel.X.(*ast.Ident); ok {
			return ast.NewIdent(ident.Name + "." + sel.Sel.Name)
		}
		return getEndIdent(sel.X)
	}
	if _, ok := expr.(*ast.MapType); ok {
		return ast.NewIdent("object")
	}
	return expr.(*ast.Ident)
}

func parseField(f *ast.Field) *StructField {
	if len(f.Names) > 0 {
		ident := getEndIdent(f.Type)
		field := NewStructField(f.Names[0].Name, ident.Name)
		if _, ok := f.Type.(*ast.StarExpr); ok {
			field.Pointer = true
		}
		if f.Doc != nil {
			for _, c := range f.Doc.List {
				field.Annotations.addLine(c.Text)
			}
		}
		if f.Comment != nil {
			for _, c := range f.Comment.List {
				field.Annotations.addLine(c.Text)
			}
		}
		return field
	}
	return nil
}

func parseStructSpec(decl *ast.GenDecl) (s *Struct) {
	if spec, ok := decl.Specs[0].(*ast.TypeSpec); ok {
		if structType, ok := spec.Type.(*ast.StructType); ok {
			if decl.Doc != nil {
				s = NewStruct(spec.Name.Name)
				for _, cg := range decl.Doc.List {
					s.Annotations.addLine(cg.Text)
				}
				for _, f := range structType.Fields.List {
					parsed := parseField(f)
					if parsed != nil {
						s.Fields = append(s.Fields, parseField(f))
					}
				}
			}
		}
	}
	return
}
