package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
                "fmt"
)

func main() {
                var a int
                Abc := 10
                ABC := 20
                ABc := 30
                
                //fmt.Printf("%v, %v", x , y)
}`

	f, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := map[string]string{}

	for _, d := range f.Decls {
		if f, ok := d.(*ast.FuncDecl); ok {
			var block = f.Body
			for _, stmt := range block.List {

				if v, ok := stmt.(*ast.AssignStmt); ok {
					if x, ok := v.Lhs[0].(*ast.Ident); ok {
						var xu = strings.ToUpper(x.Name)
						val, ex := m[xu]
						if ex {
							fmt.Printf("variable name with different casing found, old [%v] , new [%v].\n", val, x.Name)
						} else {
							m[xu] = x.Name
						}
					}
				}
			}
		}
	}
}
