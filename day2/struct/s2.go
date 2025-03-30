package main

import (
	"fmt"
	"sort"
)

func main() {
	// 定义一个string,Student的map
	m := make(map[string]Student)
	// 向m中添加元素
	m["stu01"] = Student{"tom", 18}
	m["stu02"] = Student{"mary", 20}
	m["stu03"] = Student{"jack", 19}
	// 遍历m
	for k, v := range m {
		fmt.Printf("%s:%v\n", k, v)
	}
	// 给stu01的年龄增加10岁
	m["stu01"] = m["stu01"].AddAge(10)
	temp := m["stu02"]
	s := &temp
	subAge(s)
	fmt.Printf("%v", *s)

	// m按照年龄从大到小排序
	// 1. 先将m中的元素取出，放入一个slice中
	// 2. 对slice进行排序
	// 3. 遍历slice，输出元素
	// 4. 输出结果

}

// Student 定义一个student结构体
type Student struct {
	Name string
	Age  int
}

// AddAge 函数：给student增加指定年龄
func (s Student) AddAge(age int) Student {
	s.Age += age
	return s
}

// 传入一个student，减少年龄5岁
func subAge(s *Student) {
	s.Age -= 5
}

// m按照年龄从大到小排序
// 1. 先将m中的元素取出，放入一个slice中
// 2. 对slice进行排序
// 3. 遍历slice，输出元素
// 4. 输出结果
func sortByAge(m map[string]Student) {
	// 1. 先将m中的元素取出，放入一个slice中
	slice := make([]Student, 0, len(m))
	for _, v := range m {
		slice = append(slice, v)
	}
	sort.Slice(slice, func(i, j int) bool { return slice[i].Age < slice[j].Age })
}
