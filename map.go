package main

import "fmt"

type Pos struct {
    Lat, Lng float64
}

var mos = Pos{55.7500000, 37.6155600}
var goog = Pos{37.42202, -122.08408}

var m = map[string]*Pos{
    "Moscow": &mos,
    "Google": &goog,
}

func main() {

    mos.Lat = 55.7522200
    fmt.Println(m["Moscow"])
}
