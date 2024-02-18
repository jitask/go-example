package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := []int{0, 1, 1, 2, 3, 5, 8}
	seq = slices.DeleteFunc(seq, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(seq)

	strs := []string{"asdfb", "djf", "fgjk", "fff", "ldfkg", "uiui"}
	strs = slices.DeleteFunc(strs, func(n string) bool {
		return n == "djf"
	})

	fmt.Println(strs)
}
