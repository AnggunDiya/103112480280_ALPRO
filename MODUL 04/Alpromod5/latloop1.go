package main

import "fmt"

func main(){
	var n, jumlah int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		jumlah += i
	}
	fmt.Printf("%d\n", jumlah)
}