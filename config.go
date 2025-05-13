package sorted

type RunnerConfig struct {
	All bool

	CheckConst           bool
	CheckConstSingleLine bool

	CheckVar           bool
	CheckVarSingleLine bool

	CheckStruct bool

	Report bool
}
