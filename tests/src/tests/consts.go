package tests

const (
	A1 = iota
	B1

	B2 // want `B2, A2 are not sorted alphabetically`
	A2

	A3
	B3

	B4 // want `B4, A4 are not sorted alphabetically`
	A4

	A5, B5, C5 = 0, 0, 0
	C6, B6, A6 = 0, 0, 0 // want `single line idents are not sorted alphabetically`
)

const (
	Z = iota // want `Y, X are not sorted alphabetically` `Z, Y are not sorted alphabetically`
	Y
	X

	Z1 = iota // want `Z1, X1 are not sorted alphabetically`
	X1
)

func insideFunc() {
	const (
		A1 = iota
		B1

		B2 // want `B2, A2 are not sorted alphabetically`
		A2

		A3
		B3

		B4 // want `B4, A4 are not sorted alphabetically`
		A4
	)
}
