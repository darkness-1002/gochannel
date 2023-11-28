package main

import (
	// "fmt"
	"fmt"
	"sync"
	demo "temp/demo"
)

func main() {
	// Create a buffered channel to pass data from Func1 to main
	var m sync.Mutex
	dataChannel := make(chan demo.ResultWithError, 1)

	// Create a wait group to wait for Func1 to finish
	var wg sync.WaitGroup

	
	wg.Add(1)
	// go Func1(dataChannel, &wg)
	go demo.Func3(dataChannel, &wg, &m)

	// Wait for Func3 to completes
	wg.Wait()

	// Close the channel after Func1 has finished
	close(dataChannel)
	// println(<-dataChannel)

	// Receive the result from Func1
	result, ok := <-dataChannel

	// Check if the channel was closed
	if !ok {
		fmt.Println("Channel closed in main.go.")
		return
	}
	println("reached main.go")
	// Check for errors and print the updated result
	if result.APIError != nil {
		fmt.Printf("Error in Func1: %v\n", result.APIError)
	} else {
		// Add 10 to the result and print
		updatedResult := result.RespBody + 10
		fmt.Printf("Updated result from Func1 in maing.go is : %v\n", updatedResult)
	}
}
