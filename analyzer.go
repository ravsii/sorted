package sorted

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type Runner struct {
	inspector *inspector.Inspector
	checker   *checker
}

func NewAnalyzer() *analysis.Analyzer {
	var flagSet flag.FlagSet
	_ = flagSet.Bool("test1", false, "test1")

	return &analysis.Analyzer{
		Name:     "sorted",
		Doc:      "Checks if blocks (structs, consts, vars) and functions are sorted",
		Run:      (new(Runner)).Run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    flagSet,
	}
}

func (r *Runner) Run(pass *analysis.Pass) (any, error) {
	r.inspector = pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	r.checker = newChecker(pass)

	filter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.StructType)(nil),
		// (*ast.SwitchStmt)(nil),
	}

	r.inspector.Preorder(filter, func(node ast.Node) {
		switch node := node.(type) {
		case *ast.GenDecl:
			r.validateGenDecl(pass, node)
		case *ast.SwitchStmt:
			validateSwitchStmt(pass, node)
		case *ast.StructType:
			r.validateStruct(pass, node)
		default:
			fmt.Printf("unexpected type %T\n", node)
		}
	})

	return nil, nil
}

func (r *Runner) validateStruct(pass *analysis.Pass, str *ast.StructType) {
	if !str.Struct.IsValid() {
		return
	}

	fields := str.Fields.List

	if len(fields) == 0 {
		return
	}

	nodes := make(nodes, len(fields))
	for _, f := range fields {
		nodes = append(nodes, node{
			stard: f.Pos(),
			end:   f.End(),
			Names: f.Names,
			Line:  pass.Fset.Position(f.Pos()).Line,
		})
	}

	r.checker.Check(nodes)
}

func (r *Runner) validateGenDecl(pass *analysis.Pass, decl *ast.GenDecl) {
	if !decl.Lparen.IsValid() {
		return
	}

	if decl.Tok != token.CONST && decl.Tok != token.VAR {
		return
	}

	specs := decl.Specs
	if len(specs) == 0 {
		return
	}

	nodes := make(nodes, len(specs))
	for _, spec := range specs {
		val, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		nodes = append(nodes, node{
			stard: spec.Pos(),
			end:   spec.End(),
			Names: val.Names,
			Line:  pass.Fset.Position(spec.Pos()).Line,
		})
	}

	r.checker.Check(nodes)
}

func validateSwitchStmt(pass *analysis.Pass, stmt *ast.SwitchStmt) {
	// TODO: this

	// if !stmt.Switch.IsValid() {
	// 	fmt.Println("switch statement at", stmt.Pos(), "is invalid")
	// 	return
	// }
	//
	// for _, b := range stmt.Body.List {
	// 	b := b.(*ast.CaseClause)
	// 	for _, e := range b.List {
	// 		fmt.Printf("%T", e)
	// 	}
	// }

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
