package main

import (
	"fmt"

	easy "github.com/parthvinchhi/leet-code/Easy"
)

func main() {
	// haystack := "sadbutsad"
	haystack := "varun"
	needle := "n"

	res := easy.FindtheIndex(haystack, needle)
	fmt.Println(res)
}
