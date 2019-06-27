package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPhoneContact(t *testing.T) {

	t.Run("Test phone number", func(t *testing.T) {
		phoneList := []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890", "1234567892", "(123)456-7892"}
		testResult := map[string]int{
			"1234567890": 2,
			"1234567891": 1,
			"1234567892": 3,
			"1234567893": 1,
			"1234567894": 1,
		}

		phoneBook := phoneContact(phoneList)
		fmt.Printf("%#v\n", phoneBook)
		if !reflect.DeepEqual(phoneBook, testResult) {
			t.Errorf("expect % #v but got % #v", testResult, phoneBook)
		}
	})

}
