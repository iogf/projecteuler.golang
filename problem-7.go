package main

import (
	"fmt"
)

func isPrime(n int) bool{
	for i := 2; i < n/2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}


func find(max int) int {
	count := 1
	prime := 0

	for i := 2; count <= max; i++ {
		i = i + 1
		if isPrime(i) {
			prime = i
			count = count + 1
		}

	}
	
	return prime
}


func main() {
	fmt.Println("Result:", find(10000))
}

