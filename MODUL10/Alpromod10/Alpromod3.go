package main

import "fmt"

func main() {
	var b int
	fmt.Scanln(&b)

	if b <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
	}

	fmt.Printf("Bilangan: %d\nFaktor: ", b)
	prima := true

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
			if i != 1 && i != b { 
				prima = false
			}
		}
	}

	fmt.Printf("\nPrima: %v\n", prima)
}
