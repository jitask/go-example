// Golang program to illustrate the usage of
// io.Pipe() function

// Including main package
package main

import (
	"container/ring"
	"fmt"
	"time"
)

var (
	recentMessages = ring.New(5)
	messageChan    = make(chan string)
	rereadMessage  = make(chan string)
)

func main() {
	go run()

	ticker := time.NewTicker(time.Second)
	ticker2 := time.NewTicker(3 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			messageChan <- t.Local().String()
		case <-ticker2.C:
			rereadMessage <- ""
		}
	}
}

func run() {
	for {
		select {
		case msg := <-messageChan:
			fmt.Println("init init")
			recentMessages.Value = msg
			recentMessages = recentMessages.Next()
		case <-rereadMessage:
			recentMessages.Do(func(value any) {
				fmt.Println(value)
			})
		}
	}
}
