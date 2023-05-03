package main

import (
	"fmt"
	"strings"
)

func munculSekali(str string) []string {

	runeSet := make(map[rune]struct{})
	var result []rune
	for _, r := range str {
		if _, ok := runeSet[r]; !ok {
			runeSet[r] = struct{}{}
			result = append(result, r)
		}
	}

	slice := strings.Split(string(result), " ")
	return slice
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
