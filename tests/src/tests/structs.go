package tests

type typeX struct {
	A string
	B string

	B1 string // want `B1, A1 are not sorted alphabetically`
	A1 string

	A2, B2, C2 string
	C3, B3, A3 string // want `single line idents are not sorted alphabetically`
}

type typeZ struct {
	B string // want `B, A are not sorted alphabetically`
	A string
}
