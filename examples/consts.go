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
)

const (
	Z = iota
	Y

	Z1 = iota
	Y1
)

func insideFuncConsts() {
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
