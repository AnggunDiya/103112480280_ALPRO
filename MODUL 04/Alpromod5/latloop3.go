package main

import "fmt"

func main() {
	var basis, eksponen, hasil int
	fmt.Scan(&basis, &eksponen)

	hasil = 1 
	
	for i := 1; i <= eksponen; i++ {
		hasil *= basis
	}

	fmt.Println(hasil)
}
