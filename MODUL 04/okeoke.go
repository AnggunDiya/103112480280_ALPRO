package main

import "fmt"

func main() {
	var kayu, berat, totalBerat, potong int
	fmt.Scan(&kayu)

	for i := 1; i <= kayu; i++ {
		fmt.Scan(&berat)
		totalBerat += berat
	}

	potong = (totalBerat * 9) / 10
	fmt.Println(totalBerat, potong)
}
