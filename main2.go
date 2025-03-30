package main

func main() {

	// 循环5次
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		// 打印i的值
		println("i =", 100/i)
	}

	array1 := []int{1, 2, 3, 4, 5}
	for i := range array1 {
		println(array1[i])
	}
}
