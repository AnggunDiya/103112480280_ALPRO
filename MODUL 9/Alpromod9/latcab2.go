package main

import "fmt"

func main(){
	var a int
	var teks string
	fmt.Scan(&a)
	
	if a < 0 && a%2 ==0 {
		teks = "genap negatif"
	} else {
		teks= "bukan"
	}
	fmt.Println(teks)
}