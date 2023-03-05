package main

import (
	"fmt"
)

type student struct {
	name string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode string = ""
	const offset = 3
	var toASCII []byte
	toASCII = append(toASCII, s.name...)
	
	for _, val := range toASCII{
		if (val + offset) > 122{
			val = val + offset - 26
			nameEncode += string(val)
		}else{
			val = val + offset
			nameEncode += string(val)
		}
	}
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	const offset = 3
	var toASCII []byte
	toASCII = append(toASCII, s.nameEncode...)
	
	for _, val := range toASCII{
		if (val - offset) < 97{
			val = val - offset + 26
			nameDecode += string(val)
		}else{
			val = val - offset
			nameDecode += string(val)
		}
	}
  	return nameDecode
}

func main() {
	var menu int
    var a student = student{}
	var c Chiper = &a
	

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	
	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + c.Decode())
	}
}