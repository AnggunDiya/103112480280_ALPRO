package main

import "fmt"

func main() {
	var target, donasi, totalDonasi, jumlahDonatur int
	fmt.Scan(&target)

	totalDonasi = 0
	jumlahDonatur = 0

	for totalDonasi < target {
		jumlahDonatur++
		fmt.Printf("Donatur %d: ", jumlahDonatur)
		fmt.Scan(&donasi)

		totalDonasi += donasi
		fmt.Printf("Donatur %d: Menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur\n", totalDonasi, jumlahDonatur)
}