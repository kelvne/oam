package rnp

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"

	"github.com/kelvne/oam/v2/astutils"
)

// ParseFolder parses a folder and extracts resources
func (p *Parser) ParseFolder() error {
	paths := make([]string, 0)
	if err := filepath.WalkDir(p.Path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			paths = append(paths, path)
		}
		return nil
	}); err != nil {
		return err
	}

	packages := make(map[string]*ast.Package)

	for _, path := range paths {
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}
		for k, v := range pkgs {
			packages[k] = v
		}
	}

	for _, pkg := range packages {
		for _, f := range pkg.Files {
			if err := p.parseFile(f); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *Parser) parseFile(f *ast.File) error {
	structs := astutils.ExtractResources(f)
	for _, s := range structs {
		p.Resources = append(p.Resources, resourceFromStruct(s))
	}
	return nil
}
