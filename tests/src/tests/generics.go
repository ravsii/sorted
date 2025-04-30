package tests

type (
	A string
	B string
)

type (
	GenericMap[T A | B]    map[T]bool
	GenericMapBad[T B | A] map[T]bool
)

func x[T A | B]() {}
