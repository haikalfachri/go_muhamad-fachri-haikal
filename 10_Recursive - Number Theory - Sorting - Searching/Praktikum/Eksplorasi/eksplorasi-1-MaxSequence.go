package main

import "fmt"

func MaxSequence(arr []int) int {
	var max int

	for i:=0; i<len(arr); i++{
		for j:=i+1; j<len(arr); j++{
			seq := arr[i:j]
			sum := 0
			for _, bil := range seq{
				sum += bil
			}
			if sum > max{
				max = sum
			}
		}
	}
	return max
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))  // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
  	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}