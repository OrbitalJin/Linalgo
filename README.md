# Linalgo

Linear algebra library for Go, written in Go

run test suite:

```bash
make test
```

build library:

```bash
make build
```

## Features

-   Matrices

*   [x] Addition/Substraction
*   [x] Scalar/Transformation
*   [x] Transpose
*   [x] Negate
*   [x] Value Setting
*   [x] Equality Checking
*   [x] Transposing
*   [x] Dot Product
*   [x] Identity
*   [x] SubMatrix
*   [x] Determinant (Requires SubMatrix)
*   [x] Minor
*   [x] Cofactors (Requires Minor)
*   [x] Adjugate (Requires Cofactors)
*   [x] Inverse (Requires Determinant)
*   [x] Division (A/B = A \* B^-1) (Requires Inverse)
*   [x] Row Swapping
*   [x] Row Transforation
*   [x] Gaussian Row Reduction Echelon Form (RREF) (Requires Row Swapping)
*   [x] Linear Systems
*   [x] Gauss-Jordan Elimination Solver
*   [x] Random Matrix
*   [x] Clip
*   [x] Mean
*   [x] Mul - Element wise product
*   [x] Masking (one hot encoding)
*   [x] Add tests
*   [x] Add vectors
*   [ ] Add benchmarks
