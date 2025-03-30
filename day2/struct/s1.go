package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// 用取地址的形式初始化一个student结构体
	m := &student{
		Name: "wang",
		Age:  18,
	}
	fmt.Println(*m)
	// 用构造函数初始化一个student结构体
	n := student{"wang", 18}
	fmt.Println(n)
	// 先声明一个指针，再set值
	var p *student
	p = &student{}
	p.Name = "wang"
	p.Age = 18
	fmt.Println(*p)

	// 循环初始化多个student结构体到数组中
	var arrayStudent = make([]student, 0, 10)
	for i := 0; i < 10; i++ {
		var s = student{
			// name随机字符串
			Name: generateRandomString(5),
			// age为10-20之间的随机数
			Age: rand.Intn(10) + 10,
		}
		arrayStudent = append(arrayStudent, s)
	}
	fmt.Printf("%v", arrayStudent)
}

// 定义一个student结构体
type student struct {
	Name string
	Age  int
}

// 定义一个全局的随机数生成器
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// 生成指定长度的随机字符串
func generateRandomString(length int) string {
	// 定义字符集
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		// 生成随机索引
		index := r.Intn(len(charset))
		// 根据索引从字符集中选取字符并写入字符串构建器
		sb.WriteByte(charset[index])
	}
	return sb.String()
}
