package main

import "fmt"

func main() {
	var bilangan int
	var pembagi int
	var prima bool = true

	fmt.Println("Masukkan sebuah bilangan:")
	fmt.Scanln(&bilangan)

	for i := 2; i <= bilangan/2; i++ {
		pembagi = bilangan % i
		if pembagi == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Printf("%d adalah bilangan prima\n", bilangan)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", bilangan)
	}
}
