package main

import (
	"errors"
	"fmt"
)

func main() {
	defer testRecover()
	rlt, err := testError(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rlt)
	divide(10, 0)
	fmt.Println(1234)
}

func testError(num1 int, num2 int) (rlt int, err error) {
	err = nil
	if num2 == 0 {
		err := errors.New("can't divide zero")
		return 0, err
	}
	rlt = num1 / num2
	return rlt, err
}

func divide(i int, j int) {
	defer testRecover()
	fmt.Println(i / j)
	fmt.Println("66666") // it will not be printed
}

func testRecover() {
	err := recover()
	if err != nil {
		fmt.Println("哈哈哈")
	}
}
