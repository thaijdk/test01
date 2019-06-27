package main

import (
	"fmt"
	"testing"
)

func TestRomanNumber(t *testing.T) {

	t.Run("Test roman number", func(t *testing.T) {
		var result, testResult string
		fmt.Printf("result %#v expect %#v\n", result, testResult)
		if romanNum(1) != "I" {
			t.Errorf("expect % #v", romanNum(1))
		}
		if romanNum(4) != "IV" {
			t.Errorf("expect % #v", romanNum(4))
		}
		if romanNum(5) != "V" {
			t.Errorf("expect % #v", romanNum(5))
		}
		if romanNum(9) != "IX" {
			t.Errorf("expect % #v", romanNum(9))
		}
		if romanNum(10) != "X" {
			t.Errorf("expect % #v", romanNum(10))
		}
		if romanNum(40) != "XL" {
			t.Errorf("expect % #v", romanNum(40))
		}
		if romanNum(50) != "L" {
			t.Errorf("expect % #v", romanNum(50))
		}
		if romanNum(90) != "XC" {
			t.Errorf("expect % #v", romanNum(90))
		}
		if romanNum(100) != "C" {
			t.Errorf("expect % #v", romanNum(100))
		}
	})
}
