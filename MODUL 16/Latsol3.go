package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1) // Mengatur seed untuk hasil yang konsisten

	var JmlTetes int
	fmt.Scan(&JmlTetes)

	const UkTetes = 0.0001
	hitung := [4]int{}

	for i := 0; i < JmlTetes; i++ {
		x, y := rand.Float64(), rand.Float64()
		if x < 0.5 {
			if y < 0.5 {
				hitung[2]++ // Daerah C
			} else {
				hitung[0]++ // Daerah A
			}
		} else {
			if y < 0.5 {
				hitung[3]++ // Daerah D
			} else {
				hitung[1]++ // Daerah B
			}
		}
	}

	for i, c := range hitung {
		fmt.Printf("Curah hujan daerah %c: %.4f milimeter\n", 'A'+rune(i), float64(c)*UkTetes*1000)
	}
}
