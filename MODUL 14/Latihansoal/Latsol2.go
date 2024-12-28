package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan < 2 {
		fmt.Printf("%d bukan prima\n", bilangan)
		return
	}

	statusPrima := true 

	for i := 2; i*i <= bilangan; i++ {
		if bilangan%i == 0 {
			statusPrima = false
			break 
		}
	}

	if statusPrima {
		fmt.Print("prima")
	} else {
		fmt.Print("bukan prima")
	}
}
