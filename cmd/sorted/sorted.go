package main

import (
	"github.com/ravsii/sorted"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(sorted.NewAnalyzer(nil))
}
