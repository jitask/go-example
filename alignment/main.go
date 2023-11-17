package main

import (
	"fmt"
	"unsafe"
)

type Args struct {
	a int8
	b int16
	c string
}

type Example struct {
	a int8
	b string
	c int8
	d int32
	e []string
	f Args
}

func main() {
	var f = Args{}
	var v = Example{
		a: 10,
		b: "val",
		c: 20,
		d: 100,
		e: []string{"ele"},
		f: f,
	}

	fmt.Println("Offset of a", unsafe.Offsetof(f.a)) // 0
	fmt.Println("Offset of b", unsafe.Offsetof(f.b)) // 8
	fmt.Println("Offset of c", unsafe.Offsetof(f.c)) // 16
	fmt.Println("Sizeof Example", unsafe.Sizeof(f))  // 32
	fmt.Println()
	fmt.Println("Offset of a", unsafe.Offsetof(v.a)) // 0
	fmt.Println("Offset of b", unsafe.Offsetof(v.b)) // 8
	fmt.Println("Offset of c", unsafe.Offsetof(v.c)) // 24
	fmt.Println("Offset of d", unsafe.Offsetof(v.d)) // 28
	fmt.Println("Offset of e", unsafe.Offsetof(v.e)) // 32
	fmt.Println("Offset of f", unsafe.Offsetof(v.f)) // 56
	fmt.Println("Sizeof Example", unsafe.Sizeof(v))  // 88
}
