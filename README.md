Sizeof tip for Go types
=======================

Just a small tip to check yourself when organizing structs.
```go
import . "unsafe"
```


## Fixed sizes

| Type | `Sizeof()` bytes |
| ---- | ---------------: |
| `struct{}` | 0 |
| `[0]Type` | 0 |
| `bool` | 1 |
| `int8`, `uint8`, `byte` | 1 |
| `int16`, `uint16` | 2 |
| `int32`, `uint32`, `rune` | 4 |
| `float32` | 4 |
| `int`, `uint` | 4 or 8 |
| `int64`, `uint64` | 8 |
| `float64` | 8 |
| `complex64` | 8 |
| `uintptr` | 8 |
| `*struct{}`, `*Type` | 8 |
| `map[Type1]Type2` | 8 |
| `chan Type` | 8 |
| `func()` | 8 |
| `string` | 16 |
| `complex128` | 16 |
| `[]Type` | 24 |



## Array

Formula: `Sizeof([N]Type) = N * Sizeof(Type)`

Examples:

| Type | `Sizeof()` bytes |
| ---- | ---------------: |
| `[5]bool` | 5 |
| `[2][]bool{}` | 48 |



## Struct

Struct size depends on how struct is packed.

Formula:  
`Sizeof(struct{T1; ... TN}) = StructPackSize * PacksNum`,  
where `StructPackSize = min(max(Sizeof(T1), ... Sizeof(TN)), Sizeof(uintptr))`  
and `PacksNum` depends on how struct fields are packed due to their order.

Examples:

| Type | `StructPackSize` | `PacksNum` | `Sizeof()` bytes |
| ---- | ---------------: | ---------: | ---------------: |
| `struct{a struct{}}` | 0 | 1 | 0 |
| `struct{a struct{}; b bool}` | 1 | 1 | 1 |
| `struct{b bool; u int32}` | 4 | 2 | 8 |
| `struct{a bool; b bool; u int32}` | 4 | 2 | 8 |
| `struct{u int32; a bool; b bool}` | 4 | 2 | 8 |
| `struct{a bool; u int32; b bool}` | 4 | 3 | 12 |
| `struct{a bool; s string; b bool}` | 8 | 4 | 32 |
