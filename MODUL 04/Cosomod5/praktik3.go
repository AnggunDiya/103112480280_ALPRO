package main

import "fmt"

func main() {
	var hasil, v1, v2 int
	fmt.Scan(&v1, &v2)
	hasil=0
	for j := 1; j <= v2; j+=1 {
		hasil = hasil + v1
	}
	fmt.Println(hasil)
}
