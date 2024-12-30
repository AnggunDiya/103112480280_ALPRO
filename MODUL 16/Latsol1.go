package main

import "fmt"

func main(){
    var bilangan float64
    total := 0.0
    jumlah := 0

    for {
        fmt.Scan(&bilangan)
        if bilangan == 9999 {
            break
        }
        total += bilangan
        jumlah++
    }

    if jumlah > 0 {
        rerata := total / float64(jumlah)
        fmt.Printf("%.2f\n", rerata) 
    }
}
