package main

import (
    "fmt"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    
    strLen := len(s)
    m := make(map[string]int)

    firstIdx := 0
    currWord := ""
    
    for i := 0; i < strLen; i++ {
        
        if s[i] != 32 {
            
            continue
        }
        if firstIdx == i {
            
            firstIdx++
            continue
        }
        currWord = s[firstIdx:i]
        m[currWord]++
        firstIdx = i + 1
    }
    if s[strLen - 1] != ' ' {
        currWord = s[firstIdx:strLen]
        m[currWord]++
    }
    return m
}

func main() {
    
    wc.Test(WordCount)
    fmt.Println(WordCount(" Stas   Stas stas "))
}
