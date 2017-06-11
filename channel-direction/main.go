package main

import "fmt"

func ping_func(pings chan<- string, msg string) {
	pings <- msg
}

func pong_func(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping_func(pings, "This is a message")
	pong_func(pings, pongs)

	fmt.Println(<-pongs)
}
