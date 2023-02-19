package main

import "fmt"

func main() {
	var input int
	fmt.Print("Input	: ")
	fmt.Scan(&input)
	fmt.Println("Output	: ")
	for i:=input; i>0; i--{
		for j:=1; j<i; j++{
			fmt.Print(" ")
		}
		for k:=0; k<(input+1-i); k++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}