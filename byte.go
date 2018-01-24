package main

import (
	"fmt"
)

func main() {

	var data [10]byte = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	var str1 string = string(data[:5])
	var str2 string = string(data[5:])
	var str3 string = string(data[1:3])
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
