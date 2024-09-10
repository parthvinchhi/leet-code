package input

import "fmt"

func SingleInput() int {
	var n int
	fmt.Println("Enter a number : ")
	fmt.Scanln(&n)

	return n
}

func MultipleInput() []int {
	var n []int
	var no, num int

	fmt.Println("how many input you want to enter? : ")
	fmt.Scanln(&no)

	for i := 0; i < no; i++ {
		fmt.Println("Enter a number : ")
		fmt.Scanln(&num)
		n = append(n, num)
	}

	return n
}
