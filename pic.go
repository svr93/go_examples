package main

import (
    "fmt"
    "golang.org/x/tour/pic"
)

func main() {

    pic.Show(Pic)   // base64
}

func Pic(dx, dy int) [][]uint8 {

    buf := make([][]uint8, dy)
    for i := range buf {

        buf[i] = make([]uint8, dx)
        row := buf[i]
        for j := range row {

            row[j] = uint8((i + j) / 2)
            // no overflow errors because of using variables, not constants
        }
    }
    fmt.Println(buf)
    return buf
}
