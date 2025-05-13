package tests

func insideFunc() {
	const C1, B1, A1 = 0, 0, 0 // want `single line idents are not sorted alphabetically`

	const (
		B2 = iota // want `B2, A2 are not sorted alphabetically`
		A2
	)

	// TODO: these should throw an error
	// C3, B3, A3 := 0, 0, 0 // TODO: single line idents are not sorted alphabetically
	// C3, B3, A3 = 1, 1, 1  // TODO: single line idents are not sorted alphabetically

	// this is fine
	// _ = []any{C3, B3, A3}

	var (
		B4 = 0 // want `B4, A4 are not sorted alphabetically`
		A4 = 0
	)

	_ = []any{B4, A4}

	results, err := f()
	_, _ = results, err
}

func f() (any, error) { return nil, nil }
