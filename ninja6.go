package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (r square) area() float64 {
	return r.length * r.length
}

func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println("area: ", x)
}

// CALLBACK
func useCallback(f func(arr []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func main() {
	//sq := square {
	//	length: 16.3,
	//}
	//
	//ci := circle {
	//	radius: 3.22,
	//}
	//
	//// Polimorfism, it detects whethe which struct's area method should called
	//sq.area()
	//ci.area()
	//
	//info(sq)
	//info(ci)
	//
	//res := first()
	//fmt.Println(res())

	g := func(arr []int) int {
		if len(arr) == 0 {
			return 0
		}

		if len(arr) == 1 {
			return arr[0]
		}

		return arr[0] + arr[len(arr) - 1]
	}
	val := useCallback(g, []int { 22, 33, 66 })
	fmt.Println(val)
}

