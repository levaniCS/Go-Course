package main

import "fmt"

type person struct {
 	first string
 	last string
 	flavors []string
}

func main() {
	p1 := person {
		first: "Levanin",
		last: "sarishvili",
		flavors: []string { "Vanila", "Martini" },
	}

	p2 := person{
		first: "Tornike",
		last: "Tchitadze",
		flavors: []string { "Chocolate", "Strawberry", "Capuccino" },
	}

	// Create map using struct
	m := map[string] person {
		p1.last: p1,
		p2.last: p2,
	}


	// Loop map
	for _,v := range m {

		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, val := range  v.flavors {
			fmt.Println(i, val)
		}
	}

}