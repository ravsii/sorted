package analyzer

import "go/ast"

func identsToStrings(idents []*ast.Ident) []string {
	results := make([]string, len(idents))
	for i := range idents {
		results[i] = idents[i].String()
	}

	return results
}
