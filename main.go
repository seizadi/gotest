package main

import (
	"fmt"
)


func main() {

	c := make(chan int)
	e := make(chan bool)

	// Producer
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		e <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		e <- true
	}()

	go func() {
		<- e
		<- e
		close(c)
	}()

	// Consumer
	func() {
		for i:= range c {
			fmt.Println(i)
		}
	}()
	
}