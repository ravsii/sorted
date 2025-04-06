package examples

var (
	a1 = 1
	b1 = 1

	b2 = 1
	a2 = 1

	a3 = 1
	b3 = 1

	b4 = 1
	a4 = 1
)

var (
	z = 1
	y = 1

	z1 = 1
	y1 = 1
)

func insideFuncVars() {
	var (
		a1 = 1
		b1 = 1

		b2 = 1
		a2 = 1
	)

	_ = a1
	_ = b1
	_ = b2
	_ = a2
}
