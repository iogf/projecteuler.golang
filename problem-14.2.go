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

var ctzNums map[int]int = make(map[int]int)

func CtzChain(num int) int {
	len, ok := ctzNums[num]

	if !ok {
		len = calcChain(num)
		ctzNums[num] = len
	}

	return len
}

func calcChain(num int) int {
	len := 1

	for ; num > 1; len++ {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = 3*num + 1
		}

		value, ok := ctzNums[num]

		if ok {
			return len + value
		} else {
			len++
		}
	}

	return len
}

func main() {
	num, max := 1, 1

	for i := 1; i < 1000000; i++ {
		len := CtzChain(i)

		if len > max {
			num, max = i, len
		}
	}

	fmt.Println("Num:", num)
	fmt.Println("Chain Length:", max)
}
