package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var substring string

	for i := range a{
		for j := i+1; j < len(a)+1; j++{
			tempSubstr := a[i:j]
			if strings.Contains(b, tempSubstr) && len(tempSubstr) > len(substring) {
                substring = tempSubstr
			}
		}
	}
	return substring
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI")) 		// AKA
	fmt.Println(Compare("KANGOORO", "KANG")) 	// KANG
	fmt.Println(Compare("KI", "KIJANG"))		// KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))	// KUPU
	fmt.Println(Compare("ILALANG", "ILA"))		// ILA
}