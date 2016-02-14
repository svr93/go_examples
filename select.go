/**
 * https://tour.golang.org/concurrency/6
 *
 * select statement using.
 */
package main

import (
    "fmt"
    "time"
)

func main() {

    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    
    fmt.Printf("\n'tick' type: %T", tick)
    fmt.Printf("\n'boom' type: %T", boom)
    for {

        select {

            case <-tick:
                fmt.Println("tick.")

            case <-boom:
                fmt.Println("BOOM!")
                return

            default:
                fmt.Println("    .")
                time.Sleep(50 * time.Millisecond)
        }
    }
}
