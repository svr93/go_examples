/**
 * https://tour.golang.org/concurrency/3
 *
 * Simple buffered channel using.
 */
package main

import "fmt"

func sum(a []int, ch chan int) {

    sum := 0
    for _, v := range a {

        sum += v
        ch <- sum
    }
}

func main() {

    arr := []int{7, 2, 8, -9, 4, 0}
    ch := make(chan int, len(arr))

    go sum(arr, ch)

    for i := 0; i < len(arr) / 2; i++ {

        fmt.Println(<-ch)
    }
}
