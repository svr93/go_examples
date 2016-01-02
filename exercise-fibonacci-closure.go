package main

import "fmt"

func fibonacci() func() int {

    count := 0
    prevFirst := 1
    prevSecond := 1

    return func() int {

        count++
        if count < 3 { return 1 }

        res := prevFirst + prevSecond
        prevFirst = prevSecond
        prevSecond = res

        return res
    }
}

func main() {

    f := fibonacci()
    for i := 0; i < 10; i++ {

        fmt.Println(f())
    }
}
