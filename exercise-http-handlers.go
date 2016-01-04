/**
 * https://tour.golang.org/methods/14
 *
 * Exercise: HTTP Handlers
 *
 * Implement types String, Struct and define ServeHTTP methods on them.
 * Register them to handle specific paths in your web server. 
 */
package main

import (
    "log"
    "net/http"
    "fmt"
)

type String string

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    fmt.Fprint(w, str)
}

type Struct struct {

    Greeting string
    Punct    string
    Who      string
}

func (data *Struct) String() string {

    return data.Greeting + data.Punct + " " + data.Who
}

func (data *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    fmt.Fprint(w, data)
}

func main() {
    
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

    log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
