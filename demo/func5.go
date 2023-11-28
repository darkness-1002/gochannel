// func5.go
package demo

import (
	"fmt"
	"sync"
	"time"
)

func Func5(ch3 <-chan ResultWithError, ch4 chan<- ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	defer m.Unlock()

	fmt.Println("Func5 reached")

	data, ok := <-ch3
	println("ok is ", ok)
	time.Sleep(2 * time.Second)
	if !ok {
		fmt.Println("Channel ch4 closed.")
		return
	}
	respBody := data.RespBody
	fmt.Println("responseBody is ", respBody)
	respBody *= 10
	fmt.Printf("Nihar has %d chocolates\n", respBody)

	fmt.Printf("Func5 is executing with data: %v\n", respBody)

	// Send data back to Func3
	ch4 <- ResultWithError{RespBody: respBody, APIError: nil}
}
