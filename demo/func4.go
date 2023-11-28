package demo

// package main

import (
	"fmt"
	"sync"
	"time"
)

func Func4(ch1 chan ResultWithError, ch2 chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	println("Func4 reached")
	data, ok := <-ch1

	time.Sleep(2 * time.Second)
	if !ok {
		fmt.Println("Channel closed in func4.")
		return
	}

	respBody := data.RespBody
	println("responseBody is ", respBody)
	respBody = respBody * 2

	fmt.Printf("Func4 is executing with data: %v\n", respBody)

	// Send data back to Func3
	ch2 <- ResultWithError{RespBody: respBody, APIError: nil}
	m.Unlock()
	wg.Done()
}
