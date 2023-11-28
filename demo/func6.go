package demo

// package main

import (
	"sync"
)

func Func6(ch1 chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer wg.Done()
	println("Func6 reached")
	a := 50
	b := 70
	res := a + b
	print("res in func6 is ", res)
	ch1 <- ResultWithError{RespBody: res, APIError: nil}

	m.Unlock()
}
