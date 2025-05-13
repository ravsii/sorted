package sorted

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type (
	nodes []node
	node  struct {
		stard, end token.Pos
		Names      []*ast.Ident
		Values     []ast.Expr
		Line       int
	}

	// Deprecated: use [Reporter].
	ReporterOld interface {
		Reportf(pos token.Pos, format string, args ...any)
		ReportRangef(rng analysis.Range, format string, args ...any)
	}

	Reporter func(analysis.Diagnostic)
)

type Checker interface {
	Check(nodes nodes)
}

var (
	_ Checker = (*checker)(nil)
	_ Checker = (*noOpChecker)(nil)
)

type checker struct {
	// Deprecated: use r
	reportedOld ReporterOld

	reporter *analysis.Pass
}

func newChecker(r *analysis.Pass, old ReporterOld) *checker {
	return &checker{reporter: r, reportedOld: old}
}

func (c *checker) Check(nodes nodes) {
	lastLineIdent := ""
	lastLaneNum := 0
	startedAt := token.Pos(0)

	for _, node := range nodes {
		pos := node.stard

		if startedAt == 0 {
			startedAt = pos
		}

		c.CheckSingleLine(node.Names, node.Values)

		firstIdent := ""
		if len(node.Names) > 0 {
			firstIdent = node.Names[0].String()
		}

		curLine := node.Line
		if lastLaneNum != 0 && curLine-lastLaneNum > 1 {
			lastLaneNum = curLine
			lastLineIdent = firstIdent
			startedAt = node.stard

			continue
		}

		if lastLineIdent != "" && firstIdent < lastLineIdent {
			c.reportedOld.Reportf(startedAt, `%s, %s are not sorted alphabetically`,
				lastLineIdent, firstIdent)
		}

		lastLaneNum = curLine
		lastLineIdent = firstIdent
	}
}

type declPair struct {
	name  *ast.Ident
	value ast.Expr
}

func (c *checker) CheckSingleLine(names []*ast.Ident, values []ast.Expr) {
	if len(names) < 2 {
		return
	}

	hasValues := len(values) > 0
	if hasValues && len(names) != len(values) {
		return
	}

	pairs := make([]declPair, len(names))
	for i := range names {
		pairs[i] = declPair{
			name: names[i],
		}
		if hasValues {
			pairs[i].value = values[i]
		}
	}

	// Sort by identifier name
	slices.SortFunc(pairs, func(a, b declPair) int {
		return strings.Compare(a.name.Name, b.name.Name)
	})

	// Reconstruct sorted names and values as source code
	var buf bytes.Buffer

	for i, p := range pairs {
		if i > 0 {
			buf.WriteString(", ")
		}

		buf.WriteString(p.name.Name)
	}

	if hasValues {
		buf.WriteString(" = ")

		for i, p := range pairs {
			if i > 0 {
				buf.WriteString(", ")
			}
			// Convert ast.Expr back to source
			var valBuf bytes.Buffer
			if err := format.Node(&valBuf, c.reporter.Fset, p.value); err != nil {
				return // or handle error
			}

			buf.Write(valBuf.Bytes())
		}
	}

	identStrings := identsToStrings(names)
	if !slices.IsSorted(identStrings) {
		start := names[0].Pos()

		var end token.Pos

		if hasValues {
			end = values[len(values)-1].End()
		} else {
			end = names[len(names)-1].End()
		}

		c.reporter.Report(analysis.Diagnostic{
			Pos:     start,
			End:     end,
			Message: "single line idents are not sorted alphabetically",
			SuggestedFixes: []analysis.SuggestedFix{{
				Message: "Sort declarations alphabetically",
				TextEdits: []analysis.TextEdit{{
					Pos:     start,
					End:     end,
					NewText: buf.Bytes(),
				}},
			}},
		})
	}
}

type noOpChecker struct{}

func (*noOpChecker) Check(nodes) {}
