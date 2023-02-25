package main

import "fmt"

func Mapping(slice []string) map[string] int{
	result := map[string]int{}

	for _, v := range(slice){
		result[v]++
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) //map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"})) //map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{})) //map[]
}