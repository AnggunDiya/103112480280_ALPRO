package main

import "fmt"

func main() {
    var nam float64
    var nmk string
    
    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)
    
    if nam > 88 {
        nmk = "A"
    } else if nam > 72.5 && nam <= 88 {
        nmk = "AB"
    } else if nam > 65 && nam <= 72.5 {
        nmk = "B"
    } else if nam > 57.5 && nam <= 65 {
        nmk = "BC"
    } else if nam > 50 && nam <= 57.5 {
        nmk = "C"
    } else if nam > 40 && nam <= 50 {
        nmk = "D"
    } else {
        nmk = "E"
    }
    
    fmt.Printf("Nilai mata kuliah: %s\n", nmk)
}