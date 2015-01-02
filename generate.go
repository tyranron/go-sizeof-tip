package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"unsafe"
)

type row struct {
	types []string
	sizes []uintptr
}

var FixedTable = []row{
	row{[]string{"struct{}"}, []uintptr{unsafe.Sizeof(struct{}{})}},
	row{[]string{"[0]Type"}, []uintptr{unsafe.Sizeof([0]bool{})}},
	row{[]string{"bool"}, []uintptr{unsafe.Sizeof(true)}},
	row{[]string{"int8", "uint8", "byte"}, []uintptr{unsafe.Sizeof(byte(0))}},
	row{[]string{"int16", "uint16"}, []uintptr{unsafe.Sizeof(int16(0))}},
	row{[]string{"int32", "uint32", "rune"}, []uintptr{unsafe.Sizeof(rune(0))}},
	row{[]string{"float32"}, []uintptr{unsafe.Sizeof(float32(0))}},
	row{[]string{"int", "uint"}, []uintptr{unsafe.Sizeof(int32(0)), unsafe.Sizeof(int64(0))}},
	row{[]string{"int64", "uint64"}, []uintptr{unsafe.Sizeof(int64(0))}},
	row{[]string{"float64"}, []uintptr{unsafe.Sizeof(float64(0))}},
	row{[]string{"complex64"}, []uintptr{unsafe.Sizeof(complex(float32(0), float32(0)))}},
	row{[]string{"uintptr"}, []uintptr{unsafe.Sizeof(uintptr(0))}},
	row{[]string{"*struct{}", "*Type"}, []uintptr{unsafe.Sizeof(&struct{}{})}},
	row{[]string{"map[Type1]Type2"}, []uintptr{unsafe.Sizeof(map[string]bool{})}},
	row{[]string{"chan Type"}, []uintptr{unsafe.Sizeof(make(chan bool))}},
	row{[]string{"func()"}, []uintptr{unsafe.Sizeof(func() {})}},
	row{[]string{"string"}, []uintptr{unsafe.Sizeof("")}},
	row{[]string{"complex128"}, []uintptr{unsafe.Sizeof(complex(float64(0), float64(0)))}},
	row{[]string{"[]Type"}, []uintptr{unsafe.Sizeof([]bool{})}},
}

var ArrayTable = []row{
	row{[]string{"[5]bool"}, []uintptr{unsafe.Sizeof([5]bool{})}},
	row{[]string{"[2][]bool{}"}, []uintptr{unsafe.Sizeof([2][]bool{})}},
}

var StructTable = []row{
	row{[]string{"struct{a struct{}}", "0", "1"}, []uintptr{unsafe.Sizeof(struct{ a struct{} }{})}},
	row{[]string{"struct{a struct{}; b bool}", "1", "1"}, []uintptr{unsafe.Sizeof(struct {
		a struct{}
		b bool
	}{})}},
	row{[]string{"struct{b bool; u int32}", "4", "2"}, []uintptr{unsafe.Sizeof(struct {
		b bool
		u int32
	}{})}},
	row{[]string{"struct{a bool; b bool; u int32}", "4", "2"}, []uintptr{unsafe.Sizeof(struct {
		a bool
		b bool
		u int32
	}{})}},
	row{[]string{"struct{u int32; a bool; b bool}", "4", "2"}, []uintptr{unsafe.Sizeof(struct {
		u int32
		a bool
		b bool
	}{})}},
	row{[]string{"struct{a bool; u int32; b bool}", "4", "3"}, []uintptr{unsafe.Sizeof(struct {
		a bool
		u int32
		b bool
	}{})}},
	row{[]string{"struct{a bool; s string; b bool}", "8", "4"}, []uintptr{unsafe.Sizeof(struct {
		a bool
		s string
		b bool
	}{})}},
}

func main() {
	template.Must(template.New("readme").Parse(Template)).Execute(os.Stdout,
		map[string][]string{
			"fixed":  makeTwoColumns(FixedTable),
			"array":  makeTwoColumns(ArrayTable),
			"struct": makeFourColumns(StructTable),
		},
	)
}

func makeTwoColumns(table []row) (out []string) {
	out = make([]string, len(table))
	for num, line := range table {
		sizes := make([]string, len(line.sizes))
		for i, size := range line.sizes {
			sizes[i] = fmt.Sprintf("%d", size)
		}
		out[num] = fmt.Sprintf("| `%s` | %s |\n",
			strings.Join(line.types, "`, `"),
			strings.Join(sizes, " or "),
		)
	}
	return
}

func makeFourColumns(table []row) (out []string) {
	out = make([]string, len(table))
	for num, line := range table {
		sizes := make([]string, len(line.sizes))
		for i, size := range line.sizes {
			sizes[i] = fmt.Sprintf("%d", size)
		}
		line.types[0] = "`" + line.types[0] + "`"
		out[num] = fmt.Sprintf("| %s | %s |\n",
			strings.Join(line.types, " | "),
			strings.Join(sizes, " or "),
		)
	}
	return
}

var Template = `Sizeof tip for Go types
=======================

Just a small tip to check yourself when organizing structs.
` + "```" + `go
import . "unsafe"
` + "```" + `


## Fixed sizes

| Type | ` + "`" + `Sizeof()` + "`" + ` bytes |
| ---- | ---------------: |
{{ range $row := index . "fixed" }}{{ $row }}{{ end }}


## Array

Formula: ` + "`" + `Sizeof([N]Type) = N * Sizeof(Type)` + "`" + `

Examples:

| Type | ` + "`" + `Sizeof()` + "`" + ` bytes |
| ---- | ---------------: |
{{ range $row := index . "array" }}{{ $row }}{{ end }}


## Struct

Struct size depends on how struct is packed.

Formula:  
` + "`" + `Sizeof(struct{T1; ... TN}) = StructPackSize * PacksNum` + "`" + `,  
where ` + "`" + `StructPackSize = min(max(Sizeof(T1), ... Sizeof(TN)), Sizeof(uintptr))` + "`" + `  
and ` + "`" + `PacksNum` + "`" + ` depends on how struct fields are packed due to their order.

Examples:

| Type | ` + "`" + `StructPackSize` + "`" + ` | ` + "`" + `PacksNum` + "`" + ` | ` + "`" + `Sizeof()` + "`" + ` bytes |
| ---- | ---------------: | ---------: | ---------------: |
{{ range $row := index . "struct" }}{{ $row }}{{ end }}`
