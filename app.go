package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"sort"
)

type person struct {
	first string
	age int
}

type ByAge []person

func (a ByAge) Len()int  { return len(a) }
func (a ByAge) Swap(i, j int)  { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }


func main() {

	// Unmarshall
	s := `[{"First":"Levni","Last":"Sarishvili","Age":21},{"First":"Ani","Last":"Tabatadze","Age":19}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//peoplee := []person {}
	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data:", people)

	for i,v := range people {
		fmt.Println("\nPerson number", i + 1)
		fmt.Println(v.Age, v.First, v.Last)
	}

	//Marshall
	p1:= person{
		First: "Levni",
		Last: "Sarishvili",
		Age: 21,
	}

	p2:= person{
		First: "Ani",
		Last: "Tabatadze",
		Age: 19,
	}

	people := []person { p1, p2 }
	fmt.Println("People: ", people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))


	xi := []int { 2, 3, 33, 5, 65, 102, 0, 43, 45, 65, 4 }
	xs := []string { "James", "A", "Q", "M", "Z", "z", "Money", "Dr. No", "a" }

	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)


	p1 := person{"James", 32}
	p2 := person{"Levani", 12}
	p3 := person{"Raul", 21}
	p4 := person{"Karim", 40}

	people := []person { p1, p2, p3, p4 }

	fmt.Println(people)

	// Sort by value
	sort.Sort(ByAge(people))

	fmt.Println(people)

	// Bycript
	s := "password123"

	fmt.Println(s)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	err = bcrypt.CompareHashAndPassword(bs, []byte(s))

	fmt.Println(err)


}
