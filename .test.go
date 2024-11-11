package main

import (
	"fmt"

	easy "github.com/parthvinchhi/leet-code/Easy"
)

func main() {
	haystack := "varun"
	needle := "parh"

	res := easy.FindtheIndex(haystack, needle)
	fmt.Println(res)
}
