package main

import "fmt"

func main() {
	var BMI, Tinggi_badan, Berat_badan float64 
	fmt.Scan(&BMI, &Tinggi_badan)
	Berat_badan = BMI * (Tinggi_badan * Tinggi_badan) 
	fmt.Printf("%.f", Berat_badan)
}
