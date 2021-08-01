package main

import "fmt"

type person struct {
	first string
	last string
}

func changeMe(p *person)  {
	*p = person {
		first: "Tornike",
		last: "chitadzee",
	}
}

func main() {
	v := person {
		first: "Levani",
		last: "Sarishvili",
	}

	fmt.Println("Before: ", v)

	changeMe(&v)

	fmt.Println("After: ", v)

}

