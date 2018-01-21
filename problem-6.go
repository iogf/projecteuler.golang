package main

import (
	"fmt"
	"math"
)

func main() {
	sqsum := 0
	sum := 0

	for i := 0; i <= 100; i++ {
		sqsum = sqsum + int(math.Pow(float64(i), 2))
		sum = sum + i
	}

	fmt.Println("Result:",  int(math.Pow(float64(sum), 2)) - sqsum )
}
