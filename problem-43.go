package main

import (
	"fmt"
	"strconv"
)

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func Perm(arr string) []string {
	var n func(arr []rune, p []string) []string
	n = func(arr []rune, p []string) []string {
		if len(arr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), arr[0])...)
			}
			return n(arr[1:], result)
		}
	}

	output := []rune(arr)
	return n(output[1:], []string{string(output[0])})
}

func IsStrDiv(nstr string) bool {
	p0 := []byte{nstr[1], nstr[2], nstr[3]}
	p1 := []byte{nstr[2], nstr[3], nstr[4]}
	p2 := []byte{nstr[3], nstr[4], nstr[5]}
	p3 := []byte{nstr[4], nstr[5], nstr[6]}
	p4 := []byte{nstr[5], nstr[6], nstr[7]}
	p5 := []byte{nstr[6], nstr[7], nstr[8]}
	p6 := []byte{nstr[7], nstr[8], nstr[9]}

	np0, _ := strconv.Atoi(string(p0))
	np1, _ := strconv.Atoi(string(p1))
	np2, _ := strconv.Atoi(string(p2))
	np3, _ := strconv.Atoi(string(p3))
	np4, _ := strconv.Atoi(string(p4))
	np5, _ := strconv.Atoi(string(p5))
	np6, _ := strconv.Atoi(string(p6))


	cond := (np0 % 2 == 0) && (np1 % 3 == 0) && (np2 % 5 == 0) && 
	(np3 % 7 == 0) && (np4 % 11 == 0) && (np5 % 13 == 0) &&
	(np6 % 17 == 0)

	if cond {
		return true
	}

	return false
}

func main() {
	perms := Perm("0123456789")
	var sum uint64

	for _, elem := range perms {
		if IsStrDiv(elem) {
			n, _ := strconv.Atoi(elem)
			sum += uint64(n)
			// fmt.Println(elem)
		}
	}
	fmt.Println("Sum:", sum)
}
