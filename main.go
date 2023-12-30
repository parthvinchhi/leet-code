package main

import "fmt"

func main() {
	var level string
	var num int

	fmt.Println("Let's Start!!")
	fmt.Println("Which level you want to start with : ")
	fmt.Scanln(&level)
	fmt.Println("Which number of question you want to do : ")
	fmt.Scanln(&num)

	switch {
	case level == "Easy":
		switch num {
		case 1:
			var nums []int
			var target int
			n = input.SingleInput()
			easy.TwoSum(nums, target)
		}
	}
}
