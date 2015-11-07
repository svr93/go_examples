package main

import (
    "fmt"
)

const delta = 0.0001

func Abs(x float64) float64 {

    if x > 0 { return x } else { return -x }
}

func Sqrt(x float64) float64 {
    
    z := 1.0
    prev := 0.0
    
    for Abs(z - prev) >= delta {
        prev = z
        z = z - (z * z - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
