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
	row{[]string{"bool"}, []uintptr{unsafe.Sizeof(true)}},
	row{[]string{"int8", "uint8", "byte"}, []uintptr{unsafe.Sizeof(byte(0))}},
	row{[]string{"int16", "uint16"}, []uintptr{unsafe.Sizeof(int16(0))}},
	row{[]string{"int32", "uint32"}, []uintptr{unsafe.Sizeof(int32(0))}},
	row{[]string{"int64", "uint64", "rune"}, []uintptr{unsafe.Sizeof(rune(0))}},
	row{[]string{"int"}, []uintptr{unsafe.Sizeof(int32(0)), unsafe.Sizeof(rune(0))}},
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