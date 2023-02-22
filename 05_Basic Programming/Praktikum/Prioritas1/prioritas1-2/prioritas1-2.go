package main

import "fmt"

func main() {
	var bil int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)
	if bil%2 == 0 {
		fmt.Println("Bilangan adalah genap.")
	}else{
		fmt.Println("Bilangan adalah ganjil.")
	}
}