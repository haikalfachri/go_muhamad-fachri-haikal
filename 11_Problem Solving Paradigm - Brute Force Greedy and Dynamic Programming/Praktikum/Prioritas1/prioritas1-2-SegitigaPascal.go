package main

import "fmt"

func SegitigaPascal(rows int) [][]int {
	hasil := make([][]int, rows)

	for i:=0; i<rows; i++{
		hasil[i] = make([]int, i+1)
		hasil[i][0] = 1
		hasil[i][len(hasil[i])-1] = 1
		for j:=1; j<i; j++{
			hasil[i][j] = hasil[i-1][j] + hasil[i-1][j-1]
		}
	}
	return hasil
}

func main() {
	var numRows int = 5
	fmt.Println(SegitigaPascal(numRows))
}