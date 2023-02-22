package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Angka: ")
	fmt.Scan(&angka)
	for i:=1; i<=angka; i++{
		if angka%i == 0{
			fmt.Println(i)
		}
	}
}