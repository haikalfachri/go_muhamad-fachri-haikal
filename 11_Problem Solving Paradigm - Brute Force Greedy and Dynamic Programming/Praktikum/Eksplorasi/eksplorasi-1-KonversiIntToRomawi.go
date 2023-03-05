package main

import (
    "fmt"
)

func KonversiIntToRomawi(num int) string {
    listSimbol := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
    listValue := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,}
    hasil := ""
    
    for num > 0{
        for i, val := range listValue{
            if num >= val{
                num -= val
                hasil += listSimbol[i]
            }
        }
    }
    return hasil
}

func main() {
    fmt.Println(KonversiIntToRomawi(1994)) // Output: MCMXCIV
    fmt.Println(KonversiIntToRomawi(58)) // Output: LVIII
    fmt.Println(KonversiIntToRomawi(9)) // Output: IX
    fmt.Println(KonversiIntToRomawi(4)) // Output: IV
}