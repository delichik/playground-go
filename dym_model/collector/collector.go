package main

import (
	"bytes"
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"golang.org/x/tools/imports"
)

func main() {
	var (
		pkgPath string
		outPath string
	)

	flag.StringVar(&pkgPath, "i", "", "package path")
	flag.StringVar(&outPath, "o", "./dym_gen/main.go", "output path of generator source code")
	flag.Parse()
	data, err := readPackage(pkgPath)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(outPath, data, 0644)
	if err != nil {
		panic(err)
	}
}

func readPackage(pkgPath string) ([]byte, error) {
	entries, err := os.ReadDir(pkgPath)
	if err != nil {
		return nil, err
	}

	source := bytes.NewBuffer([]byte{})
	source.WriteString("package main\n")
	source.WriteString("func main() {\n")
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.HasSuffix(entry.Name(), ".go") &&
			!strings.HasSuffix(entry.Name(), "_dym_gen.go") &&
			!strings.HasSuffix(entry.Name(), "_gen.go") &&
			!strings.HasSuffix(entry.Name(), "_test.go") {
			if err := parseFile(pkgPath, entry.Name(), source); err != nil {
				return nil, err
			}
		}
	}
	source.WriteString("}\n")

	return imports.Process("", source.Bytes(), &imports.Options{
		TabIndent: true,
	})
}

func parseFile(pkgPath string, filename string, source *bytes.Buffer) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, pkgPath+"/"+filename, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	v := &TypeSpecVisitor{
		lastTypeConfig: &TypeConfig{},
	}
	ast.Walk(v, f)

	source.WriteString("generator.Add(\"")
	source.WriteString(pkgPath)
	source.WriteString("\", \"")
	source.WriteString(strings.TrimSuffix(filename, ".go") + "_dym_gen.go")
	source.WriteString("\"")

	for _, model := range v.models {
		source.WriteString(", generator.Target{ Model: &")
		source.WriteString(model.PackageName + "." + model.StructName)
		source.WriteString("{}, TableName: \"")
		source.WriteString(model.Config.TableName)
		source.WriteString("\"}")
	}
	source.WriteString(")\n")

	return nil
}
