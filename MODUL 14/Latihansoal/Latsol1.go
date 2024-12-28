package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan <= 0 {
		fmt.Println("bilangan bulat harus positif.")
		return
	}

	jumlahGanjil := 0

	for i := 1; i <= bilangan; i++ {
		if i%2 != 0 { 
			jumlahGanjil++
		}
	}

	fmt.Printf("Terdapat %d bilangan ganjil\n", jumlahGanjil)
}

