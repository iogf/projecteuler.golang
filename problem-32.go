package main

import (
    "fmt"
    "strings"
    "strconv"
)

func Swap(arr []string, n, m int) []string {
    cp := append([]string{}, arr...)

    tmp := cp[n]
    cp[n] = cp[m]
    cp[m] = tmp
    return cp
}

func genPerms(arr []string, start int, handle func([]string)) {
    for i:=start + 1; i < len(arr); i++ {
        cp := Swap(arr, start, i)
        handle(cp)
        genPerms(cp, start + 1, handle)
    }
}

func Perm(arr []string, handle func([]string)) {
    handle(append([]string{}, arr...))
    genPerms(arr, 0, handle)
    genPerms(arr, 1, handle)
}

func isProd(x, y, z int) bool {
    if x * y == z {
        return true
    } else {
        return false
    }
}


func Slc3(arr []string, handle func(int, int, int)) {
    for i:=1; i < len(arr); i++ {
        x := arr[:i]
        y := arr[i:]

        for j:=1; j < len(y); j++ {
            z := y[:j]
            h := y[j:]

            m, _ := strconv.Atoi(strings.Join(x, ""))
            n, _ := strconv.Atoi(strings.Join(z, ""))
            k, _ := strconv.Atoi(strings.Join(h, ""))

            if isProd(m, n, k) {
                handle(m, n, k)
            }
        }
    }
}

func main() {
    handleSlcs := func(m, n, k int) {
        fmt.Println(m, n, k)
    }

    handlePerm := func(arr []string) {
        Slc3(arr, handleSlcs)
        // fmt.Println(arr)

    }

    Perm([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, handlePerm)

}

