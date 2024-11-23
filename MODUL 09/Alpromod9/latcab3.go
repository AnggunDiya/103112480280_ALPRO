package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	x_faktor_y := (y % x == 0)
	y_faktor_x := (x % y == 0)

	fmt.Println(x_faktor_y)
	fmt.Println(y_faktor_x)
}
