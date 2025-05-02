package sorted_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ravsii/sorted"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(wd, "tests")
	analysistest.Run(t, testdata, sorted.NewAnalyzer(&sorted.RunnerConfig{All: true}), "tests")
}
