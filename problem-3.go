/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"
import "math"

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
	NUM := 600851475143
	// 	NUM := 13195
	sqrt := math.Sqrt(float64(NUM))
	lim := int(sqrt) + 1
	seq := PrimeSeq(lim)

	// 	PrintSeq(seq)
	for i := len(seq) - 1; i > 0; i-- {
		if NUM%seq[i] == 0 {
			fmt.Println("Result:", seq[i])
			break
		}
	}
}
