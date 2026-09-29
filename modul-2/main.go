package main

import "fmt"

func main() {
	var nama, NIM, kelas string
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&NIM)
	fmt.Print("Masukkan kelas: ")
	fmt.Scan(&kelas)
	fmt.Println("Perkenalkan saya adalah", nama, "salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM", NIM+".")
}
