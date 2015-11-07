package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    vertex := Vertex{1, 2}

    vertexVal := vertex
    vertexVal.X = 100
    fmt.Println(vertexVal, vertex)  // {100 2} {1 2}

    vertexRef := &vertex
    vertexRef.X = 200
    fmt.Println(vertexRef, vertex)  // &{200 2} {200 2}
}
