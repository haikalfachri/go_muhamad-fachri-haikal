package main

import (
	"fmt"
)

type Car struct{
	Type string
	fuelIn float64
}

func (c Car) JarakTempuh() float64{
	return c.fuelIn / 1.5
}

func main() {
	var c Car
	fmt.Print("Masukkan Tipe Car      : ")
	fmt.Scan(&c.Type)
	fmt.Print("Masukkan Banyak Fuel   : ")
	fmt.Scan(&c.fuelIn)
	
	fmt.Println("-------------------------")
	fmt.Println("Tipe Car : ", c.Type)
	fmt.Println("Fuel     : ", c.fuelIn)
	fmt.Println("Perkiraan Jarak : ", c.JarakTempuh(), "Mil")
}