package main

import (
	"fmt" 
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) {
	var sisaMoney int = money
	var countProduct int
	
	sort.Ints(productPrice)
	for _, price := range productPrice{
		if sisaMoney - price >= 0{
			sisaMoney -= price
			countProduct++
		}
	}
	fmt.Println(countProduct)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}