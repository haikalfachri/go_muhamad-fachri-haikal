package main

import ("fmt"
		"strconv"
)

func munculSekali(angka string) []int{
	var result []int
	countAngka := map[string]int{}

	for i := range angka {
		countAngka[string(angka[i])]++
	}

	for num, count := range(countAngka){
		if count == 1{
			num, _ := strconv.Atoi(num)
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123")) // [4]
	fmt.Println(munculSekali("76523752")) // [6 3]
	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]
}