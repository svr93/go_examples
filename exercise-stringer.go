package main

import (
    "fmt"
    "strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {

    res := ""
    for _, elem := range ip {

        res += "." + strconv.Itoa(int(elem))
    }

    return res[1:]
}

func main() {

    addrs := map[string]IPAddr{

        "loopback":  { 127, 0, 0, 1 },
        "googleDNS": { 8, 8, 8, 8 },
    }
    for n, a := range addrs {

        fmt.Printf("%v: %v\n", n, a)
    }
}
