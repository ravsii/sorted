package sorted

import (
	"go/ast"
	"go/token"
	"slices"

	"golang.org/x/tools/go/analysis"
)

type (
	nodes []node
	node  struct {
		blockStart, stard, end token.Pos
		Names                  []*ast.Ident
		Line                   int
	}

	// Reported is pretty much an abstraction for [analysis.Pass] so it could
	// be easily unit-tested for integration with [analysis.Pass], not the
	// actual linter tests.
	Reporter interface {
		Reportf(pos token.Pos, format string, args ...any)
		ReportRangef(rng analysis.Range, format string, args ...any)
	}
)

type checker struct {
	r Reporter
}

func newChecker(r Reporter) *checker {
	return &checker{r: r}
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

		c.CheckNames(node.Names)
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
			c.r.Reportf(startedAt, `%s, %s are not sorted alphabetically`,
				lastLineIdent, firstIdent)
		}

		lastLaneNum = curLine
		lastLineIdent = firstIdent

	}
}

func (c *checker) CheckNames(names []*ast.Ident) {
	if len(names) < 2 {
		return
	}

	identStrings := identsToStrings(names)
	if !slices.IsSorted(identStrings) {
		iRange := newIdentRange(names[0], names[len(names)-1])
		c.r.ReportRangef(iRange, "single line idents are not sorted alphabetically")
	}
}
