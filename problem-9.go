/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
	"fmt"
	"math"
)

func genSqrts(max int) map[float64]int {
// 	lim := math.Sqrt(1000)
	sqrts := make(map[float64]int, max)

	for i := 1; i < max; i++ {
		sqrts[math.Pow(float64(i), 2)] = i
	}

	return sqrts

}

func IsPerfectSqrt(n float64) (float64, bool) {
	a := math.Sqrt(n)
	b := int(a)

	c := math.Pow(float64(b), 2)
	
	if c == n {
		return a, true
	}

	return a, false
}

func findProd(f int, sqrts map[float64]int) int{
	for i, j := range sqrts {
		for m, n := range sqrts {
			z := m + i
			v, err := IsPerfectSqrt(z)

			if  err && j + n + int(v) == f{
				return j * n * int(v)
			}
		}
	}

	return 0
	
}

func PrintSqrts(sqrts map[float64]int) {
	for key, value := range sqrts {
		fmt.Println(key, ":", value)
	}
}

func main() {
	sqrts := genSqrts(1000)
// 	PrintSqrts(sqrts)

// 	fmt.Println(IsPerfectSqrt(17))
	prod := findProd(1000, sqrts)

	fmt.Println("Result:", prod)
}



