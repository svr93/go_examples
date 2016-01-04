/**
 * https://tour.golang.org/methods/12
 *
 * Exercise: rot13Reader
 * 
 * Implement a rot13Reader that implements io.Reader
 * and reads from an io.Reader,
 * modifying the stream by applying the rot13 substitution cipher
 * to all alphabetical characters. 
 */
package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {

    r io.Reader
}

func (r13 *rot13Reader) Read(out []byte) (int, error) {

    in := make([]byte, len(out))
    n, err := r13.r.Read(in)
    if (err != nil) { return 0, err }

    for i := 0; i < n; i++ {

        if (in[i] >= 'A' && in[i] <= 'Z') {

            out[i] = getRot13Val(in[i], 'Z')
            continue
        }
        if (in[i] >= 'a' && in[i] <= 'z') {

            out[i] = getRot13Val(in[i], 'z')
            continue
        }
        out[i] = in[i]
    }
    return n, nil
}

func getRot13Val(oldVal byte, upperLimit byte) byte {

    res := oldVal + 13
    if res > upperLimit {

        return oldVal - 13
    }
    return res
}

func main() {

    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
