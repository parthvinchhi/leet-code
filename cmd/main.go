package main

import (
	"fmt"

	easy "github.com/parthvinchhi/leet-code/Easy"
	input "github.com/parthvinchhi/leet-code/Input"
)

func main() {
	var level string
	var num int

	fmt.Println("Let's Start!!")
	fmt.Println("Which level you want to start with : ")
	fmt.Scanln(&level)
	fmt.Println("Which number of question you want to do : ")
	fmt.Scanln(&num)

	switch {
	case level == "Easy" || level == "easy":
		switch num {
		case 1:
			var nums []int
			var target int
			target = input.SingleInput()
			easy.TwoSum(nums, target)
		}
	case level == "Medium" || level == "medium":
		switch num {

		}

	case level == "Hard" || level == "hard":
		switch num {
		}
	}
}
