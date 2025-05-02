# Sorted [![Build](https://github.com/ravsii/sorted/actions/workflows/build.yml/badge.svg)](https://github.com/ravsii/sorted/actions/workflows/build.yml) [![Test](https://github.com/ravsii/sorted/actions/workflows/test.yml/badge.svg)](https://github.com/ravsii/sorted/actions/workflows/test.yml)

sorted is a linter and formatter used to maintain consistent sorting across
various structures.

At the moment it only supports alphabetic sorting in `const`, `var` and
`struct` blocks, but we're aiming to support any kind of sortable or
order-related structures like `generics`, `func` in/out arguments and more.

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

These checks are working for `const`, `var`, `struct` blocks.

## TODO

- [x] Options for turning stuff on/off
  - Partially done, more options will be added later
- [x] `const`, `var`
  - [x] Alphabetical sorting
- [x] `struct`
  - [x] Alphabetical sorting
- [ ] `switch` (maybe?)
- [ ] `imports` (pretty much what gci does with some more)
- [ ] `func`
  - Alphabetical sorting
  - Public/Private sorting
  - `New*` func order
- [ ] Generics
