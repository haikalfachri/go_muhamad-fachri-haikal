package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	var x, y, z float64
	var foundSolution bool

	// x^ + y^ + z^ = C
	// xyz = B
	// x + y + z = A

	for x = 1; x*x <= float64(c); x++ { // nilai x^ || y^ || z^ <= C
		for y = x; y*y <= float64(c); y++ { 
			z = float64(a) - x - y // x + y + z = A
			if x*y*z == float64(b) && x*x + y*y + z*z == float64(c) {
				fmt.Println(int(x), int(y), int(z))
				foundSolution = true
				break
			}
		}
		if foundSolution{
			break
		}
	}
	if !foundSolution{
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3) // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}