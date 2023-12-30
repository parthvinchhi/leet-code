package input

import "fmt"

func SingleInput() int {
	var n int
	fmt.Println("Enter a number : ")
	fmt.Scanln(n)

	return n
}
