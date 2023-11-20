package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker = gobreaker.NewCircuitBreaker(gobreaker.Settings{})

func main() {
	url := "https://itask.cn"
	for {
		body, err := Get(url)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(string(body))
	}
}

func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}
