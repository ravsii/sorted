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
  const ( // or var
    B = iota // B, A are not sorted alphabetically
    A
  )
  ```

- Multiple sections treated as such

  ```go
  const ( // or var
    B = iota // B, A are not sorted alphabetically
    A

    C // ok
    D

    Z // Z, Y are not sorted alphabetically
    Y
    X
  )
  ```

- Multiple inline identifiers

  ```go
  const c, b, a = 0, 0, 0 // single line idents are not sorted alphabetically
  ```

These checks are working for `const`, `var`, `struct` blocks.

## TODO

- [x] `const`, `var`
  - [ ] Blocks
    - [x] Alphabetical sorting
    - [ ] Auto Fix
  - [ ] Inline
    - [ ] Alphabetical sorting
    - [ ] Auto Fix
- [x] `struct`
  - [x] Alphabetical sorting
  - [ ] Auto Fix
- [ ] `//sorted:ignore` comment

### TODO later

in order of most importance

- [ ] `golangci-lint` integration
- [ ] `switch` only for basic type values (string, int, ...)
- [ ] `imports` (pretty much what gci does with some more)
- [ ] `func`
  - Alphabetical sorting
  - Public/Private sorting
  - `New*` func order
- [ ] Generics
