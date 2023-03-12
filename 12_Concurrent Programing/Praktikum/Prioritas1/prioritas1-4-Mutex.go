package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	number     int
	mutex sync.Mutex
}

func (i *SafeNumber) Get() int {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	return i.number
}

func (i *SafeNumber) Set(n int) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.number = n
}

func factorial(n int) int {
	var result SafeNumber

	go func(){
		result.Set(1)
		for i:=2; i<=n; i++{
			result.Set(result.number * i)
		}
	}()
	time.Sleep(time.Second)
	return result.Get()
}

func main() {
	fmt.Println("3! adalah ", factorial(3))
	fmt.Println("5! adalah ", factorial(5))
	fmt.Println("6! adalah ", factorial(6))
}