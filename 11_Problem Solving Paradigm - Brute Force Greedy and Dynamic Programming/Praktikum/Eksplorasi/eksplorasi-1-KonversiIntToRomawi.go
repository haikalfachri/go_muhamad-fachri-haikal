package main

import (
    "fmt"
)

func KonversiIntToRomawi(num int) string {
    listSimbol := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
    listValue := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,}
    hasil := ""
    
    for i, val := range listValue{
        for num >= val{
            num -= val
            hasil += listSimbol[i]
        }
    }
    return hasil
}

func main() {
    fmt.Println(KonversiIntToRomawi(4)) // Output: IV
    fmt.Println(KonversiIntToRomawi(9)) // Output: IX
    fmt.Println(KonversiIntToRomawi(23)) // Output: XXIII
    fmt.Println(KonversiIntToRomawi(2021)) // Output: MMXXI
    fmt.Println(KonversiIntToRomawi(1646)) // Output: MDCXLVI
}