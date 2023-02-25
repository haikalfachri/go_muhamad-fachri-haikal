package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var mergedArray []string
	var isExist bool

	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(mergedArray); j++ {
			if arrayA[i] == mergedArray[j] {
				isExist = true
				break
			}
		}
		if isExist == false {
			mergedArray = append(mergedArray, arrayA[i])
		}
		isExist = false
	}

	for k := 0; k < len(arrayB); k++ {
		for l := 0; l < len(mergedArray); l++ {
			if arrayB[k] == mergedArray[l] {
				isExist = true
				break
			}
		}
		if isExist == false {
			mergedArray = append(mergedArray, arrayB[k])
		}
		isExist = false
	}
	return mergedArray

}

func main() {

	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []
}
