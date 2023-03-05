package main

import "fmt"

func primeX(number int) int{
	var isPrime bool = true
	var counter int = 0
	var n int = 2

	for counter < number{
		for i:=2; i*i <= n; i++{
			if n % i == 0{
				isPrime = false
				break
			}
		}
		if isPrime{
			counter++
		}
		if counter == number{
			return n
		}
		n++
		isPrime = true
	}
	return n
}

func main() {
	fmt.Println(primeX(1)) // 2
	fmt.Println(primeX(5)) // 11
	fmt.Println(primeX(8)) // 19
	fmt.Println(primeX(9)) // 23
	fmt.Println(primeX(10)) // 29
}