package easy

func Palindrome(x int) bool {
	originalNum := x
	reversedNum := 0

	for x > 0 {
		digit := x % 10
		reversedNum = reversedNum*10 + digit
		x /= 10
	}

	return originalNum == reversedNum
}
