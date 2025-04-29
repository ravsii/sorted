package tests

var (
	a1 = 1
	b1 = 1

	b2 = 1 // want `b2, a2 are not sorted alphabetically`
	a2 = 1

	a3 = 1
	b3 = 1

	b4 = 1 // want `b4, a4 are not sorted alphabetically`
	a4 = 1
)

var (
	z = 1 // want `z, y are not sorted alphabetically`
	y = 1

	z1 = 1 // want `z1, y1 are not sorted alphabetically`
	y1 = 1
)

func insideFuncVars() {
	var (
		a1 = 1
		b1 = 1

		b2 = 1 // want `b2, a2 are not sorted alphabetically`
		a2 = 1
	)

	_ = a1
	_ = b1
	_ = b2
	_ = a2
}
