package analyzer_test

// func TestAnalyzer(t *testing.T) {
// 	t.Parallel()
//
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		t.Fatalf("Failed to get wd: %s", err)
// 	}
//
// 	testdata := filepath.Join(wd, "tests")
//
// 	t.Run("analyze", func(t *testing.T) {
// 		t.Parallel()
// 		analysistest.Run(
// 			t, testdata, analyzer.NewAnalyzer(&sorted.RunnerConfig{All: true, Report: true}), "tests")
// 	})
//
// 	t.Run("fixes", func(t *testing.T) {
// 		t.Parallel()
// 		analysistest.RunWithSuggestedFixes(
// 			t,
// 			testdata,
// 			sorted.NewAnalyzer(&sorted.RunnerConfig{All: true, Report: false}),
// 			"fixes",
// 		)
// 	})
// }
