package main

import "fmt"

func main() {
	var r, luas float64
	fmt.Print("masukkan r = ")
	fmt.Scan(&r)
	luas = 3.14 * r * r
	fmt.Println(luas)
}
