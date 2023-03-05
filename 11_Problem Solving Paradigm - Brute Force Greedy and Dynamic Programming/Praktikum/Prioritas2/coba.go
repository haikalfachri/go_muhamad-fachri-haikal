package main

import (
	"fmt"
	"math"
)

func minCostHelper(costs [][]int, curr int, dest int, spent int, minSpent *int) {
	// jika sudah mencapai batu tujuan, update nilai minimum biaya jika biaya saat ini lebih kecil
	if curr == dest {
		if spent < *minSpent {
			*minSpent = spent
		}
		return
	}

	// jika belum mencapai batu tujuan, coba lompat ke batu berikutnya dan hitung biaya
	if curr+1 <= dest {
		newSpent := spent + int(math.Abs(float64(costs[curr][0]-costs[curr+1][0]))+math.Abs(float64(costs[curr][1]-costs[curr+1][1])))
		if newSpent < *minSpent {
			minCostHelper(costs, curr+1, dest, newSpent, minSpent)
		}
	}
	if curr+2 <= dest {
		newSpent := spent + int(math.Abs(float64(costs[curr][0]-costs[curr+2][0]))+math.Abs(float64(costs[curr][1]-costs[curr+2][1])))
		if newSpent < *minSpent {
			minCostHelper(costs, curr+2, dest, newSpent, minSpent)
		}
	}
}

func minCost(costs [][]int, n int) int {
	// inisialisasi nilai minimum biaya dengan nilai maksimum yang mungkin
	minSpent := math.MaxInt32
	// memulai dari batu 1
	curr := 0
	// mencari nilai minimum biaya dengan backtracking
	minCostHelper(costs, curr, n-1, 0, &minSpent)
	return minSpent
}

func main() {
	//map[0:[20 30] 1:[10 10] 2:[20] 3:[]]
	h := []int{10, 30, 40, 20}
	fmt.Println(minCost(h)) // output: 30
	//rute 
	//0, 1, 3
	//0, 1, 2, 3
	//0, 2, 3
	
	//map[0:[20 30] 1:[50 0] 2:[50 0] 3:[50 40] 4:[10] 5:[]]
	h = []int{30, 10, 60, 10, 60, 50}
	fmt.Println(minCost(h)) // output: 40
	//rute
	// 0, 1, 2, 3, 4, 5
	//0, 2, 3, 4, 5
	// 0, 2, 4, 5
	// 0, 2, 3, 5
}