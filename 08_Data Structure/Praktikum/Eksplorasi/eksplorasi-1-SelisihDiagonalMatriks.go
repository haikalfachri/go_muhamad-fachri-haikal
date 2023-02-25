package main

import "fmt"

func SelisihDiagonalMatriks(matriks [][]int) int{
	var d1, d2 int = 0, 0

	for baris:=0; baris<len(matriks); baris++{
		for kolom:=0; kolom<len(matriks[baris]); kolom++{
			if baris == kolom{
				d1 += matriks[baris][kolom]
			}
			if kolom == len(matriks)-1-baris{
				d2 += matriks[baris][kolom]
			}
		}
	}
	fmt.Println("d1: ", d1)
	fmt.Println("d2: ", d2)
	fmt.Printf("|%d - %d|: ", d1, d2)
	if d1 > d2{
		return d1 - d2
	}else{
		return d2 - d1
	}
}

func main() {
	var matriks [][]int = [][]int{{1, 2, 3}, 
								  {4, 5, 6}, 
								  {9, 8, 9}}
	fmt.Println(matriks)
	fmt.Println(SelisihDiagonalMatriks(matriks)) // 2)
}