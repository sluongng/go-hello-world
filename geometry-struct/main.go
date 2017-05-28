package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

type shape struct {
	start point
}

type rectangle struct {
	shape
	width int
	height int

}

type circle struct {
	shape
	radius int
}

func (r rectangle) area() int {
	return r.height * r.width
}

func main() {
	r1 := rectangle{shape{}, 10, 10}

	fmt.Println("Hello", r1.area())
}