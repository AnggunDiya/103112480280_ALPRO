package main

import "fmt"

func main() {
    var username, password string
    perc_gagal := 0

    for {
        fmt.Scan(&username, &password)

        if username == "Admin" && password == "Admin" {
            break
        }
        perc_gagal++
    }
    fmt.Println(perc_gagal, "percobaan gagal login")
}