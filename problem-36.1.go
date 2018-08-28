package main

import (
	"fmt"
	"strconv"
)

func IsPal(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	sum := 0

	for i := 0; i < 1000000; i++ {
		nstr := strconv.Itoa(i)
		if IsPal(nstr) {
			bstr := strconv.FormatInt(int64(i), 2)
			if IsPal(bstr) {
				sum += i
				fmt.Println(nstr)
			}
		}
	}

	fmt.Println("Sum:", sum)
}
