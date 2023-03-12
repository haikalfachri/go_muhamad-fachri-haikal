package main

import (
	"fmt"
	"strings"
	"regexp"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var sentence string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	var substr1 string = sentence[:len(sentence)/2]
	var substr2 string = sentence[len(sentence)/2:]
	map1, map2 := make(map[string]int), make(map[string]int)

	wg.Add(2)
	go CountLetterFreq(strings.ToLower(substr1), map1)
	go CountLetterFreq(strings.ToLower(substr2), map2)
	wg.Wait()

	CombineSolution(map1, map2)
}

func CountLetterFreq(str string, dict map[string]int){
	defer wg.Done()

	var rg, _ = regexp.Compile(`[a-z]{1}`) 
	for _, letter := range str{
		if rg.MatchString(string(letter)){
			dict[string(letter)]++
		}
	}
}

func CombineSolution(map1, map2 map[string]int){
	for key, val := range map2{
		map1[key] += val
	}

	for key, val := range map1{
		fmt.Println(key, ": ", val)
	}
}

