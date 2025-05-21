package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type Visitor struct{}

func (v *Visitor) Filter() []ast.Node {
	return []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.GenDecl)(nil),
		(*ast.StructType)(nil),
		(*ast.FuncType)(nil),
		(*ast.SwitchStmt)(nil),
	}
}

func (v *Visitor) Visit(node ast.Node) {
	switch node := node.(type) {
	case *ast.AssignStmt:
		r.validateAssignStmt(pass, node)
	case *ast.GenDecl:
		r.validateGenDecl(pass, node)
	case *ast.SwitchStmt:
		validateSwitchStmt(pass, node)
	case *ast.StructType:
		r.validateStruct(pass, node)
	case *ast.FuncType:
		r.validateFuncDecl(pass, node)
	default:
		fmt.Printf("unexpected type %T\n", node)
	}
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

func (r *Runner) validateAssignStmt(pass *analysis.Pass, decl *ast.AssignStmt) {
	specs := decl.Lhs
	if len(specs) == 0 {
		return
	}

	switch decl.Tok {
	case token.DEFINE: // :=
	case token.ASSIGN: // =
	default:
		return
	}

	nodes := make(nodes, len(specs))

	// for _, spec := range specs {
	// 	val, ok := spec.(*ast.ValueSpec)
	// 	if !ok {
	// 		continue
	// 	}
	//
	// 	nodes = append(nodes, node{
	// 		stard:  spec.Pos(),
	// 		end:    spec.End(),
	// 		Names:  val.Names,
	// 		Values: val.Values,
	// 		Line:   pass.Fset.Position(spec.Pos()).Line,
	// 	})
	// }
	//
	r.checker.Check(nodes)
}

func (r *Runner) validateGenDecl(pass *analysis.Pass, decl *ast.GenDecl) {
	if !r.genDeclShouldBeChecked(decl) {
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
			stard:  spec.Pos(),
			end:    spec.End(),
			Names:  val.Names,
			Values: val.Values,
			Line:   pass.Fset.Position(spec.Pos()).Line,
		})
	}

	r.checker.Check(nodes)
}

func (r *Runner) validateFuncDecl(_ *analysis.Pass, f *ast.FuncType) {
	if f == nil {
		return
	}

	_ = f.TypeParams
}

func (r *Runner) validateGenerics(pass *analysis.Pass, typeParams *ast.FieldList) {
	if typeParams == nil {
		return
	}

	fields := typeParams.List
	if len(fields) == 0 {
		return
	}

	nodes := make(nodes, len(fields))
	for _, field := range fields {
		nodes = append(nodes, node{
			// blockStart: decl.Pos(),
			// stard:      field.Pos(),
			// end:        field.End(),
			// Names:      val.Names,
			Line: pass.Fset.Position(field.Pos()).Line,
		})
	}

	r.checker.Check(nodes)
}

func (r *Runner) genDeclShouldBeChecked(decl *ast.GenDecl) bool {
	if decl.Tok == token.CONST && r.config.CheckConst ||
		decl.Tok == token.VAR && r.config.CheckVar {
		return true
	}

	return false
}

func validateSwitchStmt(pass *analysis.Pass, stmt *ast.SwitchStmt) {
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
