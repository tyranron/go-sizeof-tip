Sizeof tip for Go types
=======================

| Type | unsafe.Sizeof() bytes |
| ---- | --------------------- |
| `struct{}` | 0 |
| `bool` | 1 |
| `int8`, `uint8`, `byte` | 1 |
| `int16`, `uint16` | 2 |
| `int32`, `uint32`, `rune` | 4 |
| `int64`, `uint64` | 8 |
| `int`, `uint` | 4 or 8 |
| `uintptr` | 8 |
| `float32` | 4 |
| `float64` | 8 |
| `complex64` | 8 |
| `complex128` | 16 |
