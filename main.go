package main

import (
	"fmt"
)


func factorial (n int) chan int {
	out := make (chan int)
	go func() {
		result:= 1
		for i:= n; i > 0; i-- {
			result *= i
		}
		out <- result
		close(out)
	}()
	return out
}

func main() {
	c:= factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}