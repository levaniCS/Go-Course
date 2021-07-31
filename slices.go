package main

import "fmt"

func main()  {
	x := []int { 20, 10 }

	fmt.Println(x)

	// Add elements
	x = append(x, 69, 70, 71, 72)
	fmt.Println(x)

	// Merge two slices
	y := []int { 122, 123, 124, 125 }
	x = append(x, y...)
	fmt.Println(x)

	// Remove items from slice
	x = append(x[:2], x[4:]...)

	fmt.Println(x)

	jb := []string { "BMW",  "Mercedes", "Mustang", "Toyota" }
	fmt.Println(jb)

	mp := []string { "Ford",  "Opel", "Doggy", "Nissan" }
	fmt.Println(mp)

	// Multi dimensional matrix with slices
	xp := [][]string { jb, mp }
	fmt.Println(xp)

	// Maps
	m := map[string] int {
		"James": 34,
		"Ruby": 21,
	}

	// Add entry
	m["Tod"] = 33

	// Update entry
	m["James"] = 88

	// Delete entry
	delete(m, "James")

	// Check if map property actually exits
	if v, ok := m["Tod"]; ok {
		fmt.Println("THIS IS A IF PRINT: ",m["Tod"])
		fmt.Println(v)
	}

	// Loop map
	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int { 4,5,6,56 }

	// Loop slice
	for i, v := range xi {
		fmt.Println(i, v)
	}

}
