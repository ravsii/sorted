package main

import (
	"fmt"
	"go/ast"
	"go/token"

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

		if decl.Tok != token.CONST && decl.Tok != token.VAR {
			fmt.Println(decl.Tok.String(), "is not a const / var block")
			return
		}

		lastLineName := ""
		lastLaneNum := 0
		startedAt := token.Pos(0)

		for _, spec := range decl.Specs {
			s, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			pos := s.Pos()

			if startedAt == 0 {
				startedAt = pos
			}

			curLineNames := ""
			for _, name := range s.Names {
				curLineNames += name.Name
			}

			curLine := pass.Fset.Position(s.Pos()).Line
			if lastLaneNum != 0 && curLine-lastLaneNum > 1 {
				lastLaneNum = curLine
				lastLineName = curLineNames
				startedAt = s.Pos()

				continue
			}

			if lastLineName != "" && curLineNames < lastLineName {
				pass.Reportf(startedAt, "this block is not alphabetically sorted")
				pass.Reportf(s.Pos(), "here")
			}

			lastLaneNum = curLine
			lastLineName = curLineNames

		}
	})

	return nil, nil
}
