package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var kata string
	fmt.Println("Apakah Palindrome?")
	fmt.Print("masukan kata: ")
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan(){
		kata = sc.Text()
	}
	fmt.Print("captured: ", kata, "\n")
	
	var balikKata string
	for i:=len(kata)-1; i>=0; i--{
		balikKata += string(kata[i])
	}

	var isPalindrome bool = true
	for j:=range(kata){
		if (kata[j] != balikKata[j]){
			isPalindrome = false
		}
	}

	if isPalindrome{
		fmt.Println("Palindrome")
	}else{
		fmt.Println("Bukan Palindrome")
	}
}