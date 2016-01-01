package main

import (
    "math"
    "fmt"
)

func createSqrtLog(sqrt func(float64) float64) func(float64) float64 {

    return func(n float64) (res float64) {

        res = sqrt(n)
        fmt.Printf("sqrt(%f) = %f\n", n, res)
        return
    }
}

func main() {

    sqrtLog := createSqrtLog(math.Sqrt)
    sqrtLog(2)
}
