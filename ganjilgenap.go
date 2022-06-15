package ganjil_genap

import (
    "strconv"
)

func CekBilangan(bil ...int) string {

	result := ""
	
	for _, i := range bil {

		result += strconv.Itoa(i)+ ": "

		if i%2 == 0 {
			result += "genap"
		}else{
			result += "ganjil"
		}

		result += "\n"
		
	}

	return result
}

func JustTest() string {
	return "Test"
}