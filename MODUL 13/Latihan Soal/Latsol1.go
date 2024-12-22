package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	digit := 0
	if bilangan == 0 {
		digit = 1
	} else {
		for bilangan > 0 {
			bilangan /= 10
			digit++
		}
	}

	fmt.Println(digit)
}