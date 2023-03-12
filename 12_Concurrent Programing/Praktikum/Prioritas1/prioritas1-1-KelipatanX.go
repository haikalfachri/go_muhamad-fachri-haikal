package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	go KelipatanX(x)
	time.Sleep(time.Duration(x) * time.Second + time.Second)
	fmt.Println("Done!")
}

func KelipatanX(x int){
	for i:=x; ; i++{
		if i%x == 0{
			fmt.Println(i)
			time.Sleep(1 * time.Second)
			fmt.Println("1 Sec...")
			time.Sleep(1 * time.Second)
			fmt.Println("2 Sec...")
			time.Sleep(1 * time.Second)
			fmt.Println("3 Sec...")
		}
	}
}