package main

import (
	"fmt"
	"strings"
)

func main() {
	n := int64(123456789)

	str := NumToBHex(n)
	fmt.Println(str)

	v := BHexToNum(str)
	fmt.Println(v)
}

var (
	num2Char = "0123456789abcdefghijklmnopqrstuvwxyz"
	scale    = int64(36)
)

func NumToBHex(num int64) string {
	str := ""
	for num > 0 {
		mod := num % scale
		str = string(num2Char[mod]) + str
		num = num / scale
	}

	length := 8 - len(str)
	for i := 0; i < length; i++ {
		str = "0" + str
	}

	return strings.ToUpper(str)
}

func BHexToNum(str string) int64 {
	str = strings.ToLower(str)
	length := len(str)
	val := int64(0)
	for i, s := range str {
		index := strings.Index(num2Char, string(s))
		times := length - i - 1
		pow := int64(1)
		for j := 0; j < times; j++ {
			pow *= scale
		}

		fmt.Println(pow)
		val += int64(index) * pow
	}

	return val
}
