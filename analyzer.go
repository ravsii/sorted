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
		(*ast.SwitchStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.GenDecl:
			validateGenDecl(pass, n)
		case *ast.SwitchStmt:
			validateSwitchStmt(pass, n)
		default:
			fmt.Printf("unexpected type %T\n", n)
		}
	})

	return nil, nil
}

func validateGenDecl(pass *analysis.Pass, decl *ast.GenDecl) {
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
}

func validateSwitchStmt(pass *analysis.Pass, stmt *ast.SwitchStmt) {
	if !stmt.Switch.IsValid() {
		fmt.Println("switch statement at", stmt.Pos(), "is invalid")
		return
	}

	for _, b := range stmt.Body.List {
		b := b.(*ast.CaseClause)
		for _, e := range b.List {
			fmt.Printf("%T", e)
		}
	}

	// lastLineName := ""
	// lastLaneNum := 0
	// startedAt := token.Pos(0)
	//
	// for _, spec := range decl.Specs {
	// 	s, ok := spec.(*ast.ValueSpec)
	// 	if !ok {
	// 		continue
	// 	}
	//
	// 	pos := s.Pos()
	//
	// 	if startedAt == 0 {
	// 		startedAt = pos
	// 	}
	//
	// 	curLineNames := ""
	// 	for _, name := range s.Names {
	// 		curLineNames += name.Name
	// 	}
	//
	// 	curLine := pass.Fset.Position(s.Pos()).Line
	// 	if lastLaneNum != 0 && curLine-lastLaneNum > 1 {
	// 		lastLaneNum = curLine
	// 		lastLineName = curLineNames
	// 		startedAt = s.Pos()
	//
	// 		continue
	// 	}
	//
	// 	if lastLineName != "" && curLineNames < lastLineName {
	// 		pass.Reportf(startedAt, "this block is not alphabetically sorted")
	// 		pass.Reportf(s.Pos(), "here")
	// 	}
	//
	// 	lastLaneNum = curLine
	// 	lastLineName = curLineNames
	//
	// }
}
