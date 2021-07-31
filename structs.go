package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main()  {

	tr := truck {
		fourWheel: true,
		vehicle: vehicle{
			doors: 8,
			color: "red",
		},
	}

	se := sedan {
		luxury: false,
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
	}

	fmt.Println("Truck doors: ", tr.color)
	fmt.Println("Sedan doors: ",se.doors)
}