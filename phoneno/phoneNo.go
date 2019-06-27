package main

import (
	"fmt"
	"strings"
)

func phoneContact(phoneList []string) map[string]int {
	phoneBook := make(map[string]int)
	for _, phoneNum := range phoneList {
		phoneNum = strings.ReplaceAll(phoneNum, " ", "")
		phoneNum = strings.ReplaceAll(phoneNum, "-", "")
		phoneNum = strings.ReplaceAll(phoneNum, "(", "")
		phoneNum = strings.ReplaceAll(phoneNum, ")", "")
		phoneBook[phoneNum]++
	}
	return phoneBook
}

func main() {
	phoneList := []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890", "1234567892", "(123)456-7892"}
	phoneBook := phoneContact(phoneList)
	fmt.Printf("%#v\n", phoneBook)

}
