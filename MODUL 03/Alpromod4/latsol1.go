package main

import "fmt"

func main() {
	var Belanja_awal, diskon, total int 
	fmt.Scan(&Belanja_awal, &diskon)

	total = Belanja_awal - (Belanja_awal*diskon)/100 
	fmt.Println(total)
}