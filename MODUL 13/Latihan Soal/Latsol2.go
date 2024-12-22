package main

import (
	"fmt"
	"math"
)

func main() {
	var bilangan float64
	fmt.Scan(&bilangan)

	pembulatan := math.Ceil(bilangan)
	
	i := bilangan
	for i <= pembulatan {
		fmt.Printf("%.1f\n", i)
		i += 0.1
	}
}