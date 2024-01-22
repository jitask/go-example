package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < 20; i++ {
		a := rd.Intn(10)
		fmt.Println(a)
	}
}
