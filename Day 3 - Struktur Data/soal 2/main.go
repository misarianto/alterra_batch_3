package main

import "fmt"

func Mapping(slice []string) map[string]int {
	dataMap := map[string]int{}

	for _, element := range slice {
		word := element
		dataMap[word] = len(element)
	}

	return dataMap
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{}))
}
