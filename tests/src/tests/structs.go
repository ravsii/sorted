package tests

type (
	type1 struct{}
	type2 struct{ a string }
	type3 struct{ a, b string }
	type4 struct{ b, a string } // want `single line idents are not sorted alphabetically`

	type5 struct {
		a, b string
		d, c string // want `single line idents are not sorted alphabetically`
	}
	type6 struct {
		b string // want `b, a are not sorted alphabetically`
		a string
	}
)
