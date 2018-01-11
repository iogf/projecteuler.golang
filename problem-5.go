package main

import "math"
import "fmt"

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

func exhaust(base int, num int) int{
	exp := 0
	for ; num % base == 0; exp++ {
		num = num / base
	}
	return exp
}

func main() {
	lcm := 1
	seq := PrimeSeq(20)
	hash := make(map[int]int)

	for _, n := range seq {
		for i := 2; i <= 20; i++ {
			exp := exhaust(n, i)
			if hash[n] < exp {
				hash[n] = exp
			}
		}
	}

	for base, exp := range hash {
		val := math.Pow(float64(base), float64(exp))
		lcm *=  int(val)

	}

	fmt.Println("Result:", lcm)
}


