package main

import "fmt"

func main() {
	var sisiAtas, sisiBawah, tinggi float64

	fmt.Println("Masukkan panjang sisi atas trapesium:")
	fmt.Scanln(&sisiAtas)

	fmt.Println("Masukkan panjang sisi bawah trapesium:")
	fmt.Scanln(&sisiBawah)

	fmt.Println("Masukkan tinggi trapesium:")
	fmt.Scanln(&tinggi)

	luas := 0.5 * (sisiAtas + sisiBawah) * tinggi

	fmt.Println("Luas trapesium adalah:", luas)
}
