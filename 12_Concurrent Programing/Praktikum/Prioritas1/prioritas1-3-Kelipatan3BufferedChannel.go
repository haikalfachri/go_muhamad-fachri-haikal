package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int
	fmt.Print("Masukkan batas: ")
	fmt.Scan(&x)

	// buffered channel
	ch := make(chan int, x)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer close(ch)
		defer wg.Done()
		for i := 1; i <= x; i++ {
			if i % 3 == 0 {
				ch <- i
			}
		}
	}()

	fmt.Println("----Buffered Channel----")
	for i := range ch {
		fmt.Println("Receive from buffered channel: ", i)
	}
	defer fmt.Println("Done!")

	wg.Wait()
}