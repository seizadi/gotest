package main

import "fmt"


func main() {

	c := make(chan int)
	e := make(chan bool)
	// Producer
	go func() {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// Consumer
	go func() {
		for i:= range c {
			fmt.Println(i)
		}
		e <- true
	}()

	<-e
}