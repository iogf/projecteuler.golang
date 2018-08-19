/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

*/

package main

import (
	"fmt"
	"strconv"
)

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func Perm(arr string) []string {
	var n func(arr []rune, p []string) []string
	n = func(arr []rune, p []string) []string {
		if len(arr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), arr[0])...)
			}
			return n(arr[1:], result)
		}
	}

	output := []rune(arr)
	return n(output[1:], []string{string(output[0])})
}

func isProd(x, y, z int) bool {
	if x*y == z {
		return true
	} else {
		return false
	}
}

func isUniqueProd(arr []int, item int) bool {
	for _, elem := range arr {
		if item == elem {
			return false
		}
	}
	return true
}

func Slc3(arr string, handle func(int, int, int)) {
	for i := 1; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			x := arr[0:i]
			z := arr[i:j]
			h := arr[j:]

			m, _ := strconv.Atoi(x)
			n, _ := strconv.Atoi(z)
			k, _ := strconv.Atoi(h)

			if isProd(m, n, k) {
				handle(m, n, k)
			}
		}
	}
}

func main() {
	prods := make([]int, 0)
	sum := 0
	handleSlcs := func(m, n, k int) {
		if isUniqueProd(prods, k) {
			sum += k
			prods = append(prods, k)
			fmt.Println(m, n, k)

		}

	}
	perms := Perm("123456789")

	for _, elem := range perms {
		Slc3(elem, handleSlcs)
	}

	fmt.Println("The sum:", sum)
}
