package main

import (
	"fmt"
	"math"
)

func main() {
	var volumebola, luasbola, jejari float64
	fmt.Scan(&jejari)
	volumebola = 4.0 / 3.0 * math.Pi * math.Pow(jejari, 3)
	luasbola = 4 * math.Pi * math.Pow(jejari, 2)
	fmt.Printf("bola dengan jejari %.f memiliki volume %.4f dan luas kulit %.4f\n", jejari, volumebola, luasbola)
}
