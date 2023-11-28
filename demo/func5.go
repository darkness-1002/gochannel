// func5.go
package demo

import (
	"fmt"
	"sync"
	"time"
)

func Func5(ch2 chan ResultWithError, ch3 chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	// m.Lock()
	// defer m.Unlock()

	fmt.Println("Func5 reached")
	// ch3 <- ResultWithError{1, nil}
	// Receive data from Func3
	data, ok := <-ch2
	println("data recived in func5 is %v", data.RespBody)

	println("value of ok is ", ok)
	time.Sleep(2 * time.Second)
	if !ok {
		fmt.Println("Channel closed in func5.")
		return
	}
	res := data.RespBody
	println("data recived in func5 is %v", res)
	res += 10
	ch3 <- ResultWithError{res, nil}

}

// func Func5(data ResultWithError, ch2 chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {
// 	defer wg.Done()

// 	m.Lock()
// 	defer m.Unlock()

// 	fmt.Println("Func5 reached")
// 	data.RespBody = data.RespBody + 50

// 	// Send data to Func5
// 	ch2 <- data
// }
