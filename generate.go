package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type row struct {
	types []string
	sizes []uintptr
}

var Table = []row{
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
	row{[]string{"*struct{}", "*int", "*Type"}, []uintptr{unsafe.Sizeof(&struct{}{})}},
	row{[]string{"map[Type1]Type2"}, []uintptr{unsafe.Sizeof(map[string]bool{})}},
	row{[]string{"chan Type"}, []uintptr{unsafe.Sizeof(make(chan bool))}},
	row{[]string{"func()"}, []uintptr{unsafe.Sizeof(func(){})}},
	row{[]string{"string"}, []uintptr{unsafe.Sizeof("")}},
	row{[]string{"complex128"}, []uintptr{unsafe.Sizeof(complex(float64(0), float64(0)))}},
	row{[]string{"[]Type"}, []uintptr{unsafe.Sizeof([]bool{})}},
}

func main() {
	for _, line := range []string{
		"Sizeof tip for Go types",
		"=======================",
		"",
		"| Type | unsafe.Sizeof() bytes |",
		"| ---- | --------------------- |",
	} {
		fmt.Println(line)
	}
	for _, line := range Table {
		sizes := make([]string, len(line.sizes))
		for i, size := range line.sizes {
			sizes[i] = fmt.Sprintf("%d", size)
		}
		fmt.Printf("| `%s` | %s |\n",
			strings.Join(line.types, "`, `"),
			strings.Join(sizes, " or "),
		)
	}
}
