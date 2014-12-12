Sizeof tip for Go types
=======================

| Type | unsafe.Sizeof() bytes |
| ---- | --------------------- |
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
| `*struct{}`, `*int`, `*Type` | 8 |
| `map[Type1]Type2` | 8 |
| `chan Type` | 8 |
| `func()` | 8 |
| `string` | 16 |
| `complex128` | 16 |
| `[]Type` | 24 |
