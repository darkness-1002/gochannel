package demo

// package main

import (
	"fmt"
	"sync"
	// ""
)

type ResultWithError struct {
	RespBody int
	APIError error
}

func Func1(ch chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	a := 5
	b := 10
	res := a + b
	fmt.Printf("Result in Func1: %v\n", res)

	respBody, _ := Func2(res)
	ch <- ResultWithError{RespBody: respBody, APIError: nil}

	fmt.Printf("Result in Func1: %v\n", res)

	// Send the result to the channel

}
