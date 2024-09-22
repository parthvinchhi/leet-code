package easy

func PlusOne(digits []int) []int {
	var res int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		s := digits[i]
		if s+res <= 9 {
			digits[i] += res
			break
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
