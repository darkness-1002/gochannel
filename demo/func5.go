// func5.go
package demo

import (
	"fmt"
	"sync"
)

// func Func5(ch3 chan ResultWithError, ch4 chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
// 	defer wg.Done()

// 	m.Lock()
// 	defer m.Unlock()

// 	fmt.Println("Func5 reached")

// 	// Receive data from Func3
// 	// data, ok := <-ch3
// 	// println("value of ok is ", ok)
// 	// time.Sleep(2 * time.Second)
// 	// if !ok {
// 	// 	fmt.Println("Channel closed in func5.")
// 	// 	return
// 	// }
// 	// res := data.RespBody
// 	// println("data recived in func5 is %v", res)

// }

func Func5(data ResultWithError, ch3 chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	defer m.Unlock()

	fmt.Println("Func5 reached")
	data.RespBody = 100

	// Send data to Func5
	ch3 <- data
}
