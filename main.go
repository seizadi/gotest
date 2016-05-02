package main

import (
	"fmt"
	"sync"
)


func main() {

	c := make(chan int)
	e := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)
	// Producer
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
		e <- true
	}()

	// Consumer
	go func() {
		for i:= range c {
			fmt.Println(i)
		}
	}()



	<- e
	// Sometimes program closes and last character is not printed
	// time.Sleep(time.Second) will show the character
}