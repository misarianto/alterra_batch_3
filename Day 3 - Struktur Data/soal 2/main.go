package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
	dataMap := map[string]int{}

	for _, element := range slice {
		word := element
		_, matched := dataMap[word]
		if matched {
			dataMap[word] = dataMap[word] + 1
		} else {
			dataMap[word] = 1
		}
	}

	return dataMap
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{}))
}
