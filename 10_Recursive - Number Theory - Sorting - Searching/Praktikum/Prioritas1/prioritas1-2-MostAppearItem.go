package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name string
	count int
}

func MostAppearItem(items []string) []pair {
	var countItem map[string]int = map[string]int{}

	for _, item := range items{
		countItem[item]++
	}

	var item pair
	var arrItem []pair

	for k, v := range countItem{
		item.name = k
		item.count = v
		arrItem = append(arrItem, item)
	}

	sort.Slice(arrItem, func(i, j int) bool {
		return arrItem[i].count < arrItem[j].count
	  })

	return arrItem
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1, basketball->1, tenis->1
}