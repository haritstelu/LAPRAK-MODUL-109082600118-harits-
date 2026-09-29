package main

import "fmt"

func main() {
	var fahrenheit, celcius float64
	fmt.Print("masukkan suhu dalam Fahrenheit = ")
	fmt.Scan(&fahrenheit)
	celcius = (fahrenheit - 32) * 5 / 9
	fmt.Println(celcius)
}
