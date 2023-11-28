// func3.go
package demo

import (
	"fmt"
	"sync"
)

func Func3(ch1 chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	a := 5
	b := 10
	res := a + b
	fmt.Printf("Result in Func3: %v\n", res)

	// Send data to Func4
	ch1 <- ResultWithError{RespBody: res, APIError: nil}

	// Create a new channel for receiving data from Func4

	//always define size of the channel
	ch2 := make(chan ResultWithError, 1)

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go Func4(ch1, ch2, &wg2, m)

	// Receive data from Func4
	result, ok := <-ch2
	wg2.Wait()
	if !ok {
		fmt.Println("Channel closed in func3.")
		return
	}
	respBody := result.RespBody
	fmt.Println("Result in func3 after func4 is ", respBody)

	ch2 <- result
	ch3 := make(chan ResultWithError)
	wg2.Add(1)
	go Func5(ch2, ch3, &wg2, m)
	// go Func5(result, ch3, &wg2, m)

	result2, ok := <-ch3

	if !ok {
		fmt.Println("Channel closed in func3.")
		return
	}
	wg2.Wait()
	respBody2 := result2.RespBody
	fmt.Println("Result in func3 after func6 is ", respBody2)
	ch1 <- result2

	close(ch2)
	// close(ch4)
}
