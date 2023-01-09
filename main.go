package main

import "fmt"

func main() {
	for _, v := range []rune("") {
		i := int64(v)
		fmt.Println(i)
	}
}
