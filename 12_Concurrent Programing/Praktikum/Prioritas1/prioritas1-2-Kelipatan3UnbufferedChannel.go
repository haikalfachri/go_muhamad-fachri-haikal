package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int
	fmt.Print("Masukkan batas: ")
	fmt.Scan(&x)

	// Unbuffered channel
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		wg.Done()
		for i := 1; i <= x; i++ {
			if i % 3 == 0 {
				ch <- i
			}
		}
		close(ch)
	}()

	fmt.Println("------Unbuffered Channel-------")
	for i := range ch {
		fmt.Println("Receive from unbuffered channel: ", i)
	}
	defer fmt.Println("Done!")

	wg.Wait()
}