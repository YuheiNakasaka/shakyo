package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func(SendChannel chan<- int) {
		fmt.Println("0")
		close(SendChannel)
	}(ch)
	for {
		ok := <-ch
		fmt.Printf("%d\n", ok)
		if ok == 0 {
			break
		}
	}
}
