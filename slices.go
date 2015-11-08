package main

import "fmt"

func main() {
    
    var arr [10]int
    printArr(arr)
    
    s := make([]int, 10)
    printSlice(s)
}

func printArr(arr [10]int) { // need use length for param
    
    fmt.Println(arr)
}

func printSlice(s []int) {
    
    fmt.Println(s)
}
