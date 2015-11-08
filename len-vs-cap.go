package main

import "fmt"

func main() {

    b := make([]int, 0, 5)
    // b[0] = 1 ---> panic: runtime error: index out of range
    printSlice("b", b)
    
    b = b[:cap(b)]
    b[0] = 1 // ok
    printSlice("b", b)
    
    c := b[:cap(b)]
    c[1] = 2
    fmt.Println("See modifying:")
    printSlice("b", b)
    printSlice("c", c)
}

func printSlice(s string, x []int) {

    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
