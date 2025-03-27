package main

import (
	"fmt"
	"strconv"
)

/*Given two non-negative integers num1 and num2 represented as strings,
return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert
the inputs to integer directly.
*/

func multiply(num1 string, num2 string) string {
	if num2 == "0" || num1 == "0" {
		return "0"
	}
	result := ""
	for i2 := len(num2) - 1; i2 >= 0; i2-- {
		curr := ""
		carry := 0
		for i1 := len(num1) - 1; i1 >= 0; i1-- {
			mulResult := int((num1[i1]-'0')*(num2[i2]-'0')) + carry
			curr = strconv.Itoa(mulResult%10) + curr
			carry = mulResult / 10
		}

		if carry != 0 {
			curr = strconv.Itoa(carry) + curr
		}
		for i := 1; i <= len(num2)-i2-1; i++ {
			curr += "0"
		}
		result = add(result, curr)
	}
	return result
}

func add(string1 string, string2 string) string {
	result := ""
	carry := 0
	if len(string1) > len(string2) {
		string1, string2 = string2, string1
	}

	margin := len(string2) - len(string1)
	for i := 0; i < margin; i++ {
		string1 = "0" + string1
	}

	for i := len(string2) - 1; i >= 0; i-- {
		addResult := int(string1[i]-'0'+string2[i]-'0') + carry
		carry = addResult / 10
		result = strconv.Itoa(addResult%10) + result
	}

	if carry != 0 {
		result = strconv.Itoa(carry) + result
	}
	return result
}
func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}
