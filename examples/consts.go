package examples

const (
	A1 = iota
	B1

	B2
	A2

	A3
	B3

	B4
	A4

	A5, B5, C5 = 0, 0, 0
	C6, B6, A6 = 0, 0, 0
)

const (
	Z = iota
	Y

	Z1 = iota
	Y1
)

func insideFunc() {
	const (
		A1 = iota
		B1

		B2
		A2

		A3
		B3

		B4
		A4
	)
}
