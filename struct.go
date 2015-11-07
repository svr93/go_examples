package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    vertex := Vertex{10, 2}
    fmt.Println(vertex.X)
}
