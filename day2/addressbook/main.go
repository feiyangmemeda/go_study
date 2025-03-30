package main

import (
	"day1/day2/addressbook/entity"
	"fmt"
)

var personList []entity.Person

func main() {
	scanNum()
}
func scanNum() {
	for {
		fmt.Println("添加联系人信息，请按1")
		fmt.Println("删除联系人信息，请按2")
		fmt.Println("查询联系人信息，请按3")
		fmt.Println("修改联系人信息，请按4")
		var num int
		fmt.Scan(&num)
		switchType(num)
		fmt.Println("返回上一级请按0,退出请按Q")
		var exit string
		fmt.Scan(&exit)
		if exit == "Q" {
			break
		} else {
			continue
		}
	}

}

func switchType(num int) {
	switch num {
	case 1:
		// 添加联系人信息
		addPerson()
	case 2:
		// 删除联系人信息
		deletePerson()
	case 3:
		// 查询联系人信息
		listPerson()
	case 4:
		// 修改联系人信息
	}
}

func addPerson() {
	var name string
	var phoneAddress string
	fmt.Println("请输入姓名")
	fmt.Scan(&name)
	fmt.Println("请输入电话")
	fmt.Scan(&phoneAddress)
	person := entity.Person{Username: name, PhoneAddress: phoneAddress}
	personList = append(personList, person)
}

func listPerson() {
	// 如果联系人列表为空,则提示没有联系人
	if len(personList) == 0 {
		fmt.Println("没有联系人")
		return
	}
	for _, person := range personList {
		fmt.Println(person.Username, person.PhoneAddress)
	}
}

func deletePerson() {
	var name string
	fmt.Println("请输入姓名")
	fmt.Scan(&name)
	// 遍历联系人列表,找到姓名相同的联系人,删除
	for i, person := range personList {
		if person.Username == name {
			personList = append(personList[:i], personList[i+1:]...)
			return
		}
	}
}
