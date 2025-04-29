package tests

func insideFunc() {
	const C1, B1, A1 = 0, 0, 0 // want `single line idents are not sorted alphabetically`

	const (
		B2 = iota // want `B2, A2 are not sorted alphabetically`
		A2
	)

	var C3, B3, A3 = 0, 0, 0 // want `single line idents are not sorted alphabetically`

	_ = []any{C3, B3, A3}

	var (
		B4 = 0 // want `B4, A4 are not sorted alphabetically`
		A4 = 0
	)

	_ = []any{B4, A4}
}
