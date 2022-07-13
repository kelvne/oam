package astutils

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"testing"
)

func TestJSONAnnotations(t *testing.T) {
	paths := make([]string, 0)

	filepath.WalkDir("./testassets", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			paths = append(paths, "./"+path)
		}
		return nil
	})

	packages := make(map[string]*ast.Package)

	for _, path := range paths {
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
		if err != nil {
			t.Fatal(err.Error())
		}

		for k, v := range pkgs {
			packages[k] = v
		}
	}

	t.Run("ExtractResources should extract json annotations", func(t *testing.T) {
		structs := make(map[string]*Struct)

		for _, pkg := range packages {
			for _, file := range pkg.Files {
				for key, s := range ExtractResources(file) {
					structs[key] = s
				}
			}
		}

		jsonAnnotated, ok := structs["JSONAnnotatedModel"]
		if !ok {
			t.Fatal("JSONAnnotatedModel not extracted properly")
		}

		if _, ok := structs["NestedModel"]; !ok {
			t.Fatal("NestedModel not extracted properly")
		}

		rest := jsonAnnotated.Annotations.Multiple("rest")

		found := make([]bool, 0)

		for _, r := range rest {
			rp := r.(map[string]interface{})

			if rp["path"] == "/mngr/jsonannotatedmodel" || rp["path"] == "/admin/base/{baseId}/json_annotated_models" || rp["path"] == "/admin/base/{baseId}/relations/{relationId}/json_annotated_models" {
				found = append(found, true)
			}
		}

		if len(found) != 3 {
			t.Fatal("@rest annotations failed to be extracted")
		}
	})
}
