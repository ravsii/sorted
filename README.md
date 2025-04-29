# Sorted [![Build](https://github.com/ravsii/sorted/actions/workflows/build.yml/badge.svg)](https://github.com/ravsii/sorted/actions/workflows/build.yml) [![Test](https://github.com/ravsii/sorted/actions/workflows/test.yml/badge.svg)](https://github.com/ravsii/sorted/actions/workflows/test.yml)

`sorted` is the linter for keeping everything sorted.

At the moment it only checks for a few things, with plans for checking
everything that could be checked for any sort of ordering.

## So what's working?

It can be generalized into 2 main categories for now, that are

- Blocks

  ```go
  const (
      B = iota // B, A are not sorted alphabetically
      A
  )
  ```

- Multiple inline identifiers

  ```go
  const c, b, a = 0, 0, 0 // single line idents are not sorted alphabetically
  ```

## TODO

- [ ] Options for turning stuff on/off
- [x] `const`, `var`
  - Alphabetical sorting
- [x] `struct`
  - Alphabetical sorting
- [ ] `switch` (maybe?)
- [ ] `imports` (pretty much what gci does with some more)
- [ ] `func`
  - Alphabetical sorting
  - Public/Private sorting
- [ ] Generics
