// func6.go
package demo

import (
	"fmt"
	"sync"
	"time"
)

func Func6(ch chan ResultWithError, wg *sync.WaitGroup, m *sync.Mutex) {

	defer wg.Done()
	println("Func6 reached")

	// Simulating some processing in Func6
	time.Sleep(1 * time.Second)

	// Performing some computation
	result := 20

	// Print the result in Func6
	fmt.Printf("Result in Func6: %v\n", result)

	// Send the result back to the calling function
	ch <- ResultWithError{RespBody: result, APIError: nil}
}
