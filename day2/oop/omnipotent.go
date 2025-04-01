package main

import "fmt"

// omnipotent 万能的,this demo show omnipotent interface,it can input any type param
func main() {
	b := &Boy{
		Name: "zhangsan",
		Age:  18,
	}
	sayHello(b)
}

func sayHello(i interface{}) {
	fmt.Println("hello")
	_, isString := i.(string)
	fmt.Println(isString)
}

type Boy struct {
	Name string
	Age  int
}
