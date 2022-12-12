package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	roman := []string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	in := []int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	num := 0

	for i := len(roman) - 1; i >= 0; i-- {

		if strings.Contains(s, roman[i]) && strings.Index(s, roman[i]) == 0 {
			fmt.Println(roman[i], in[i])

			num += in[i]
			s = strings.Replace(s, roman[i], "", 1)
		}
	}
	return num
}
