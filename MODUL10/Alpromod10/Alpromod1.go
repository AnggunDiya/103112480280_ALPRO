package main

import "fmt"

func main() {
	var berat, sisa int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	parsel := berat / 1000
	berat_after := berat % 1000
	biaya := parsel * 10000

	if berat_after < 500 {
		sisa = berat_after * 15
	} else {
		sisa = berat_after * 5
	}

	fmt.Printf("Detail berat: %d kg + %d gr\n", parsel, berat_after)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biaya, sisa)
	fmt.Printf("Total biaya: Rp. %d\n", biaya+sisa)
}