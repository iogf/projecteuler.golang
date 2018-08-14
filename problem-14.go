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

    for i:=n;i>1; len++ {
        if i%2 == 0 {
            i = i / 2
        } else {
            i = 3 * i + 1
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

    for i:=1000000; i>1; i-- {
        max = MaxChain(max, CalcCtz(i))
    }

    fmt.Println("Num:", max.Num);
    fmt.Println("Chain Length:", max.Len);

}

