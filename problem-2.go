/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values
do not exceed four million, find the sum of the even-valued terms.
*/

package main

import "fmt"

func fibonacci(handle func(int), max int) {
	a := 1
	b := 2
	handle(a)
	handle(b)

	for b < max {
		tmp := b
		b = a + b
		a = tmp
		handle(b)
	}

}

func main() {
	sum := 0

	acc := func(n int) {
		if n%2 == 0 {
			sum = sum + n
		}
		fmt.Println(n)
	}

	fibonacci(acc, 4000000)
	fmt.Println("Result:", sum)
}
