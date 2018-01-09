/*
A palindromic number reads the same both ways. The largest palindrome made 
from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/

package main

import (
	"fmt"
	"strconv"
)

func IsPalim(num int) bool {
	nstr := strconv.Itoa(num)
	nlen := len(nstr)

	for i := 0; i <= nlen/2; i++ {
		if nstr[nlen - i - 1] != nstr[i] {
			return false
		}
	}

	return true
}

func gNumber(min, max int) int {
	lnum := 0
	for i := max; i >= min; i-- {
		for j := max; j >= min; j-- {
			num := i * j
			if IsPalim(num) {
				if num > lnum {
					lnum = num
				}
			}
		}
	}
	return lnum
}

func main() {
	fmt.Println("Result:", gNumber(100, 999))

}
