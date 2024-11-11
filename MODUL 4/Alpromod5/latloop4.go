package main

import "fmt"


func main(){
	var n, hasilFaktorial int
	fmt.Scan(&n)
	hasilFaktorial = 1

	for i := 1; i <= n; i++ {
		hasilFaktorial *= i
	}
	fmt.Println(hasilFaktorial)
}