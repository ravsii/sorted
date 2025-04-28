package sorted

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var _ analysis.Range = (*identRange)(nil)

type identRange struct {
	pos, end token.Pos
}

func newIdentRange(first, last *ast.Ident) *identRange {
	return &identRange{pos: first.Pos(), end: last.End()}
}

func (i *identRange) Pos() token.Pos { return i.pos }
func (i *identRange) End() token.Pos { return i.end }
