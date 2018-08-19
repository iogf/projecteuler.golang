/*
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var primeSeq []int

func InitializeSeq(max int) []int {
	primeSeq = PrimeSeq(max)
	return primeSeq
}

func Rotations(n int) []int {
	nstr := strconv.Itoa(n)
	seq := make([]int, 0)

	for i:=0; i < len(nstr); i++ {
		nstr = strings.Join([]string{nstr[1:], nstr[:1]}, "")
		m, _ := strconv.Atoi(nstr)
		seq = append(seq, m)
	}

	return seq
}

func IsPrime(n int) bool {
	for _, elem := range primeSeq {
		if elem == n {
			return true
		}
	}
	return false
}

func feedSeq(seq []int, num int) []int {
	sqrt := math.Sqrt(float64(num))
	lim := int(sqrt + 1)

	for _, elem := range seq {
		if num%elem == 0 {
			return seq
		}

		if elem > lim {
			break
		}
	}

	seq = append(seq, num)
	return seq
}

func PrimeSeq(max int) []int {
	seq := make([]int, 2, 100)
	seq[0] = 2
	seq[1] = 3

	for i := 5; i < max; i += 2 {
		seq = feedSeq(seq, i)
	}

	return seq
}

func IsCircularPrime(n int) bool {
	rots := Rotations(n)

	for _, elem := range rots {
		if !IsPrime(elem)  {
			return false
		}
	}

	return true
}

func main() {
	counter := 0
	seq := InitializeSeq(1000000)

	for _, elem := range seq {
		if IsCircularPrime(elem) {
			counter += 1
		}
	}
	fmt.Println("Result:", counter)
}

