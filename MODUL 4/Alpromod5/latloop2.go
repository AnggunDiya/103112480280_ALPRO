package main

import ("fmt"
"math"
)

func main() {
	var n int
	var volume, jarijari, tinggikerucut float64
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&jarijari, &tinggikerucut)
		volume = (1.0/3.0) * (math.Pi * jarijari* jarijari * tinggikerucut) 
		fmt.Println(volume)
	}
}