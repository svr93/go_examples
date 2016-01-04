/**
 * https://tour.golang.org/methods/13
 *
 * Simple http-server.
 */
package main

import (
    "fmt"
    "log"
    "net/http"
)

type reqHandler struct{}

func (rh *reqHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    fmt.Fprint(w, "Success!")
}

func main() {

    err := http.ListenAndServe("localhost:7000", &reqHandler{})
    if err != nil {

        log.Fatal(err)
    }
}
