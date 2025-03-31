package main

import "fmt"

func main() {
	// extends obj,it has special initialize
	t := teacher{
		Person: Person{
			id:   1,
			name: "wang",
			age:  18,
		},
		salary: 10000,
	}
	fmt.Printf("%v\n", t)

	s := student{
		Person: &Person{
			id:   1,
			name: "wang",
			age:  18,
		},
		score: 100,
	}
	fmt.Printf("%v\n", s)

	fmt.Printf(t.Person.name + "\n")
	fmt.Printf("%f\n", s.score)
	fmt.Printf("%d\n", s.age)
	s.age = 99
	fmt.Printf("%d\n", s.age)
	s.oid = 10086
	fmt.Println(s.oid)
	fmt.Println(s.getName())
	s.setName("段")
	fmt.Println(s.getName())

	var s2 = student{}
	s3 := student{}
	fmt.Println(s2, s3)
}

type Person struct {
	object
	id   int
	name string
	age  int
}

type teacher struct {
	Person // this is an anonymous field, it will be inherited(java extends)
	salary float64
}

type student struct {
	*Person // this is an anonymous pointer,it will be inherited(java extends)
	score   float64
}

type object struct {
	oid int
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) setName(name string) {
	p.name = name
}
