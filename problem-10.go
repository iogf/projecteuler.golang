package main

import (
	"fmt"
	"math"
)

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

func PrintSeq(seq []int) {
	for _, elem := range seq {
		fmt.Println("Elem:", elem)
	}
}

func main() {
	arr := PrimeSeq(2000000)
	sum := 0

	for _, elem := range arr {
		sum = sum + elem
	}

	fmt.Println("Result:", sum)
}


