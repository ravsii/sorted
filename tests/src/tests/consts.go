package tests

const A1 = 0

const (
	A2 = iota
	B2
)

const (
	B3 = iota // want `B3, A3 are not sorted alphabetically`
	A3
)

const (
	A4 = iota
	B4

	B5 // want `B5, A5 are not sorted alphabetically`
	A5
)

const (
	A6, B6, C6 = 0, 0, 0
	C7, B7, A7 = 0, 0, 0 // want `single line idents are not sorted alphabetically`
)

const A8, B8, C8 = 0, 0, 0

const C9, B9, A9 = 0, 0, 0 // want `single line idents are not sorted alphabetically`
