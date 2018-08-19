package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var primeSeq []int

func InitializeSeq(max int) []int {
	primeSeq = PrimeSeq(1000000)
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
	seq := InitializeSeq(300)

	for _, elem := range seq {
		if IsCircularPrime(elem) {
			counter += 1
		}
	}
	fmt.Println("Result:", counter)
}

