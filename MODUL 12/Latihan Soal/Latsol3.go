package main

import "fmt"

func main() {
    var x, y, hasil int
    fmt.Scan(&x, &y)

    if x < y {
        fmt.Println(0)
        return
    }

    for x >= y {
        x -= y
        hasil++
    }

    fmt.Println(hasil)
}