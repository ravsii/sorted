package sorted

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type (
	nodes []node
	node  struct {
		stard, end token.Pos
		Names      []*ast.Ident
		Line       int
	}

	// Reported is pretty much an abstraction for [analysis.Pass] so it could be
	// easily unit-tested for integration with [analysis.Pass], not the actual
	// linter tests.
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
	lastLineName := ""
	lastLaneNum := 0
	startedAt := token.Pos(0)

	for _, node := range nodes {
		pos := node.stard

		if startedAt == 0 {
			startedAt = pos
		}

		curLineNames := ""
		for _, name := range node.Names {
			curLineNames += name.Name
		}

		curLine := node.Line
		if lastLaneNum != 0 && curLine-lastLaneNum > 1 {
			lastLaneNum = curLine
			lastLineName = curLineNames
			startedAt = node.stard

			continue
		}

		if lastLineName != "" && curLineNames < lastLineName {
			c.r.Reportf(startedAt, "this block is not alphabetically sorted")
			c.r.Reportf(node.stard, "here")
		}

		lastLaneNum = curLine
		lastLineName = curLineNames

	}
}
