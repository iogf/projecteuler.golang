/* The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million
*/

package main

import (
	"fmt"
)

type CtzNum struct {
	Num int
	Len int
}

func CalcCtz(n int) CtzNum {
	len := 1

	for i := n; i > 1; len++ {
		if i%2 == 0 {
			i = i / 2
		} else {
			i = 3*i + 1
		}
	}

	return CtzNum{n, len}
}

func MaxChain(n CtzNum, m CtzNum) CtzNum {
	if n.Len > m.Len {
		return n
	} else {
		return m
	}
}

func main() {
	max := CalcCtz(13)

	for i := 1000000; i > 1; i-- {
		max = MaxChain(max, CalcCtz(i))
	}

	fmt.Println("Num:", max.Num)
	fmt.Println("Chain Length:", max.Len)

}
