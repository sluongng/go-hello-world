package main

import (
	"runtime"
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0

	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	go say("world")
	say("hello")

	a := []int{1, 2, 7, -22, 12, 55}

	c := make(chan int, 2) // buffer channel which store 2 ints

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <- c, <- c

	fmt.Println("X is", x)
	fmt.Println("Y is", y)
	fmt.Println("Sum is", x + y)
}