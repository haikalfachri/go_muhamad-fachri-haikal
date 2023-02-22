package main

import (
	"fmt"
)

func main() {
	var a, b, t float32
	fmt.Print("Masukkan panjang atas  : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan panjang bawah : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan tinggi	: ")
	fmt.Scan(&t)
	var luasTrapesium float32 = (a + b) * t / 2
	fmt.Println("Luas Trapesium: ", luasTrapesium)
}