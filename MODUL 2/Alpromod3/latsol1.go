package main

import "fmt"

func main() {
	var x, operasi float64

	fmt.Scan(&x)

	operasi = (2/(x+5) + 5)
	fmt.Println("Hasil dari ", x, "adalah", operasi)
}
