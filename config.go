package sorted

import "flag"

type RunnerConfig struct {
	CheckConst           bool
	CheckConstSingleLine bool

	CheckVar           bool
	CheckVarSingleLine bool

	CheckStruct bool
}

func initConfigFromFlags() (flag.FlagSet, *RunnerConfig) {
	var (
		flagSet flag.FlagSet
		config  RunnerConfig
	)

	flagSet.Init("sorted", flag.ExitOnError)

	flagSet.BoolVar(&config.CheckConst, "check-const", false, "Check const() blocks")
	flagSet.BoolVar(&config.CheckConstSingleLine, "check-const-single-line", false, "Check const blocks for multiple identifiers in a single line")

	flagSet.BoolVar(&config.CheckVar, "check-var", false, "Check var() blocks")
	flagSet.BoolVar(&config.CheckVarSingleLine, "check-var-single-line", false, "Check var blocks for multiple identifiers in a single line")

	flagSet.BoolVar(&config.CheckStruct, "check-struct", false, "Check struct field order")

	return flagSet, &config
}
