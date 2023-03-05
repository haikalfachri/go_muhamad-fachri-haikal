package main

import (
	"fmt"
	"math"

)

func Frog(jumps []int) int {
	var minCosts []int  = make([]int, len(jumps))

	for i := 1; i < len(jumps); i++ {
		// cost = cost minimal dari jalur yang sudah dilalui + |cost batu yang sedang dipijak - cost batu index sebelumnya|
		costJump1 := minCosts[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		costJump2 := 99999
		if i > 1 {
			costJump2 = minCosts[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		}
		if costJump1 < costJump2{
			minCosts[i] = costJump1
		}else{
			minCosts[i] = costJump2
		}
	}
	return minCosts[len(jumps)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}