package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

func main() {
	filename := filepath.Join(".", "./user/user.go")
	astFile, err := parser.ParseFile(&token.FileSet{}, filename, nil, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		panic(err)
	}

	parseGenDecls(astFile)
}

// GenDecls means general declarations. Example: const, var outside of functions
func parseGenDecls(file *ast.File) {
	for _, decl := range file.Decls {
		genDecl := decl.(*ast.GenDecl)
		for _, spec := range genDecl.Specs {
			typeSpec := spec.(*ast.TypeSpec)
			name := typeSpec.Name.Name
			fmt.Println(name)

			structType := typeSpec.Type.(*ast.StructType)
			if structType.Fields.List != nil {
				for _, field := range structType.Fields.List {
					fmt.Println(field.Names[0].Name, field.Type.(*ast.Ident).Name)
				}
			}
		}
	}
}
