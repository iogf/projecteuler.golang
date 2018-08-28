package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IsPal(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

func Revstr(str string) string {
	chars := make([]byte, 0)

	for i := len(str) - 1; i >= 0; i-- {
		chars = append(chars, str[i])
	}

	return string(chars)
}

func GenIntPals(max int, handle func(int)) {
	var spal string
	var npal int

	for i := 0; i < 10; i++ {
		handle(i)
	}

	for i := 1; ; i++ {
		str0 := strconv.Itoa(i)
		str1 := Revstr(str0)

		spal = strings.Join([]string{str0, str1}, "")
		npal, _ = strconv.Atoi(spal)

		if npal >= max {
			return
		}
		handle(npal)

		for j := 0; j < 10; j++ {
			str2 := strconv.Itoa(j)
			arr := []string{str0, str2, str1}
			spal = strings.Join(arr, "")
			npal, _ = strconv.Atoi(spal)

			if npal < max {
				handle(npal)
			}
		}
	}
}

func main() {
	sum := 0

	handle := func(num int) {
		bstr := strconv.FormatInt(int64(num), 2)

		if IsPal(bstr) {
			sum += num
			fmt.Println(num)
		}
	}

	GenIntPals(1000000, handle)

	fmt.Println("Sum:", sum)
}
