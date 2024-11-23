package main

import "fmt"

func main() {
	var tahunkabisat int
	var hasil bool
	fmt.Scan(&tahunkabisat)
	hasil = tahunkabisat%400 == 0 || tahunkabisat%4 == 0 && tahunkabisat%100 != 0
	fmt.Println(hasil)
}
