package main

import "fmt"

func main() {
	var bilangan int

	fmt.Println("Masukkan sebuah bilangan:")
	fmt.Scanln(&bilangan)

	if bilangan%7 == 0 {
		fmt.Printf("%d adalah kelipatan 7\n", bilangan)
	} else {
		fmt.Printf("%d bukan kelipatan 7\n", bilangan)
	}
}
