package main 

import (
	"fmt"
	"math"
)

func main() {
	var Ax, Ay, Bx, By, Cx, Cy float64 
	fmt.Scan(&Ax, &Ay)
	fmt.Scan(&Bx, &By)
	fmt.Scan(&Cx, &Cy)
	
	panjangAB := math.Sqrt(math.Pow(Bx-Ax, 2)+math.Pow(By-Ay,2)) 
	panjangBC := math.Sqrt(math.Pow(Cx-Bx, 2)+math.Pow(Cy-By,2)) 
	panjangCA := math.Sqrt(math.Pow(Ax-Cx, 2)+math.Pow(Ay-Cy,2))
	panjangterpanjang := math.Max(panjangAB, math.Max(panjangBC, panjangCA))
	
	fmt.Println(panjangterpanjang)
}
	