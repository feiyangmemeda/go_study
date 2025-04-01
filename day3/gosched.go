package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 500; i++ {
			fmt.Println("it's goroutine func")
		}

	}()

	for i := 0; i < 500; i++ {
		// leave the current cpu time slice
		runtime.Gosched()
		fmt.Println("it's main func")
	}

}
