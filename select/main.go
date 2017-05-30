package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1

	// this loops will run forever until `return` is called
	for {
		// select helps listening to many channel
		//
		// select is blocking by default
		// and will only continue when one channel has data
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			// here return is called when there is some data pushed into channel `quit`
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			break
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
