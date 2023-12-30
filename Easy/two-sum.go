func twoSum(nums []int, target int) []int {
	var sum int
	var i int
	var output []int

	for i, sum = range nums {
		sum += nums[i]
		if sum < target {
			output = []int{i}
		}
	}
	return output
} 