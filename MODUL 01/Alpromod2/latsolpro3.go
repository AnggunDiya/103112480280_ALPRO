package main

import (
	"fmt"
	"math"
)

func main() {
	var r, luas float64
	fmt.Scan(&r)
	luas = math.Pi * math.Pow(r, 2)

	fmt.Printf("%.1f\n", luas)
}
