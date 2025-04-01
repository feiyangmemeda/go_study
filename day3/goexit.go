package main

import (
	"fmt"
	"runtime"
)

// which difference between return and Goexit?

func main() {
	fmt.Println("start ... ")

	go func() {
		fmt.Println("aaaaa")
		test2()
		defer fmt.Println("bbbbb")

	}()

	last := runtime.GOMAXPROCS(5)
	fmt.Println(last)
	for {

	}

}

func test() {
	defer fmt.Println("ccccc")
	return
	fmt.Println("ddddd")
}

func test2() {
	defer fmt.Println("ccccc")
	runtime.Goexit()
	fmt.Println("ddddd")
}
