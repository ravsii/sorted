package p

func notPrintfFuncAtAll() {}

func funcWithEllipsis(args ...interface{}) {}

func printfLikeButWithStrings(format string, args ...string) {}

func printfLikeButWithBadFormat(format int, args ...interface{}) {}

func prinfLikeFunc(format string, args ...interface{}) {} // want "printf-like formatting function"

func prinfLikeFuncWithReturnValue(format string, args ...interface{}) string { // want "printf-like formatting function"
	return ""
}
