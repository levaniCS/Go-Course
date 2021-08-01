package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

// Empty interface: interface {} - every value implements empty interface

// Method of Secret Agent
func (s secretAgent) speak() {
	fmt.Println("Hello i'm,", s.first, " - the secretAgent speak")
}

// Method of Person
func (s person) speak() {
	fmt.Println("Hello i'm,", s.first, " - the person speak")
}

// interfaces: 'if you have this methods you are also my type'
type human interface {
	speak()
}

// Differenciate using interface
func bar(h human) {
	switch h.(type) {
		case person:
			fmt.Println("I was passed into bar", h.(person).first)
		case secretAgent:
			fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
}

func main()  {
	sa1 := secretAgent {
		ltk: true,
		person: person {
			first: "James",
			last: "Bond",
		},
	}

	sa2 := secretAgent {
		ltk: true,
		person: person {
			first: "Miss",
			last: "Money",
		},
	}


	p1 := person {
		first: "Levani",
		last: "Sarishvili",
	}

	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	// Anonymous functions
	func(msg string) {
		fmt.Println("Test Anonymous:",msg)
	}("Message")

	//	Function expression
	f := func() {
		fmt.Println("Hello i am function expression")
	}
	f()


	testCopy := test()()

	fmt.Printf("%T\n", testCopy)
}

func test() func() int  {
	return func() int {
		return 420
	}
}