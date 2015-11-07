package main

import "fmt"

func main() {

    const BIG = 1 << 100
    fmt.Println(BIG >> 99)  // 2
}
