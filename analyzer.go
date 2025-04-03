package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"slices"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var analyzer = &analysis.Analyzer{
	Name:     "goprintffuncname",
	Doc:      "Checks that printf-like functions are named with `f` at the end.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.GenDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		decl := node.(*ast.GenDecl)

		if !decl.Lparen.IsValid() {
			fmt.Println(decl.Tok.String(), "is invalid")
			return
		}

		if decl.Tok != token.CONST {
			fmt.Println(decl.Tok.String(), "is not a const block")
			return
		}

		data := []string{}
		lastLine := 0
		startAt := token.Pos(0)

		reset := func() {
			data = []string{}
			lastLine = 0
			startAt = token.Pos(0)
		}

		for _, spec := range decl.Specs {
			s, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			pos := s.Pos()

			curLine := pass.Fset.Position(s.Pos()).Line
			if lastLine != 0 && curLine-lastLine > 1 {
				if !slices.IsSorted(data) {
					pass.Reportf(startAt, "this block is not alphabetically sorted")
				}

				reset()
			}
			lastLine = curLine

			if startAt == 0 {
				startAt = pos
			}

			names := ""
			for _, name := range s.Names {
				names += name.Name
			}
			data = append(data, names)

		}

		if len(data) > 0 {
			if slices.IsSorted(data) {
				return
			}

			pass.Reportf(startAt, "this block is not alphabetically sorted")
		}
	})

	return nil, nil
}
