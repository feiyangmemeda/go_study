package main

import (
	"fmt"
	"time"
)

func main() {

	go sing()
	go dance()

	func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i am goroutine")
			time.Sleep(time.Second)
		}
	}()

	// forever blocked
	select {}

}

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("you are my destiny~~...")
		time.Sleep(time.Millisecond * 100)
	}

}

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("dongci daci~~...")
		time.Sleep(time.Millisecond * 100)
	}

}
