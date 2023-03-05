package main

import "fmt"

func RepresentasiBiner(n int) []string{
	var ans []string
	for i:=0; i<=n; i++{
		ans = append(ans, ConvertIntToBinary(i))
	}
	return ans
}

func ConvertIntToBinary(n int) string {
    if n == 0 {
        return "0"
    }else if n == 1 {
        return "1"
    }

    if n % 2 == 0 {
        return ConvertIntToBinary(n / 2) + "0"
    }else {
        return ConvertIntToBinary(n / 2) + "1"
    }
}

func main() {
	fmt.Println("2: ", RepresentasiBiner(2)) // [0, 1, 10]
	fmt.Println("3: ", RepresentasiBiner(3)) // [0, 1, 10, 11]
}