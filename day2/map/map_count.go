package main

import "fmt"

func main() {
	var content = "hello world"
	var contentMap = make(map[string]int)
	for _, char := range content {
		strChar := string(char)
		num := contentMap[strChar]
		if num == 0 {
			contentMap[strChar] = 1
		} else {
			contentMap[strChar] = num + 1
		}
		// 高阶写法
		// contentMap[strChar]++
	}
	fmt.Println(contentMap)

	var content2 = "你好世界"

	for _, char := range content2 {
		fmt.Printf("%c", char)
	}
}
