package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	rj := person{
		name: "RJ",
		age:  32,
	}

	rj.age = 50
	fmt.Println(rj)
	fmt.Println(rj.name, "is a new guy")
	fmt.Println(rj.name, "is", rj.age, "years old")

	// fmt.Println(person{"Bob", 20})

	// fmt.Println(person{"Alice", 30})

	// fmt.Println(person{name: "Fred"})

	// fmt.Println(&person{name: "Ann", age: 40})

	// s := person{name: "Sean", age: 50}

	// sp := &s

	// fmt.Println(sp.age)

	// sp.age = 51
	// fmt.Println(sp.age)

}
