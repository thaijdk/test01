package main

import "fmt"

func romanNum(input int) string {
	result := ""
	romanNum := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanString := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for index := 0; index < len(romanNum); index++ {
		for romanNum[index] <= input {
			input = input - romanNum[index]
			result = result + romanString[index]
		}
	}
	return result
}

func main() {
	fmt.Printf("%#v\n", romanNum(1))
	fmt.Printf("%#v\n", romanNum(4))
	fmt.Printf("%#v\n", romanNum(5))
	fmt.Printf("%#v\n", romanNum(9))
	fmt.Printf("%#v\n", romanNum(10))
	fmt.Printf("%#v\n", romanNum(40))
	fmt.Printf("%#v\n", romanNum(50))
	fmt.Printf("%#v\n", romanNum(90))
	fmt.Printf("%#v\n", romanNum(100))
}
