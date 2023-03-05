package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	var toASCII []byte
	toASCII = append(toASCII, input...)
	var toHuruf string

	offset = offset % 26
	for _, val := range toASCII{
		if (val + byte(offset)) > 122{
			new_val := val + byte(offset) - 26 
			toHuruf += string(new_val)
		}else{
			new_val := val + byte(offset)
			toHuruf += string(new_val)
		}
	}
	return toHuruf
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
  	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}