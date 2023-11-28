package demo

// package main

import (
	"fmt"
	"time"
)

// Func2 simulates an API call and returns a response body and an error

func Func2(data int) (int, error) {
	// Simulate some processing in Func2
	time.Sleep(2 * time.Second)
	data *= 2

	// Simulate an API response
	respBody := data

	fmt.Printf("Func2 is executing with data: %v\n", data)

	return respBody, nil
}

// func Func1(ch chan<- ResultWithError, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	a := 5
// 	b := 10
// 	res := a + b
// 	fmt.Printf("Result in Func1: %v\n", res)

// 	// Call Func2 and pass the result to the channel
// 	respBody, err := Func2(res)
// 	if err != nil {
// 		ch <- ResultWithError{APIError: err}
// 		return
// 	}
// 	fmt.Printf("Result in Func1: %v\n", res)

// 	// Send the result to the channel
// 	ch <- ResultWithError{RespBody: respBody, APIError: nil}
// }

// func Func2(ch chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()

// 	// Simulate some processing in Func2
// 	// time.Sleep(2 * time.Second)
// 	println("Func2 reached")
// 	data, ok := <-ch

// 	println(ok)
// 	time.Sleep(2 * time.Second)
// 	if !ok {
// 		fmt.Println("Channel closed.")
// 		return
// 	}

// 	// Simulate an API response
// 	respBody := data.RespBody
// 	respBody *= 2

// 	fmt.Printf("Func2 is executing with data: %v\n", data)
// 	ch <- ResultWithError{RespBody: respBody, APIError: nil}
// 	m.Unlock()
// 	wg.Done()

// 	// return respBody, nil
// }
