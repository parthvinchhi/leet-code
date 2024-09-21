package medium

import (
	"math"
	"sort"
)

// {1,1,1,2,2,3}
func RemoveDuplicates(nums []int) int {
	var count, returnableCount int
	var num = nums[0]
	// fmt.Println(num)
	for key, val := range nums[1:] {
		// fmt.Println(val)
		if val == num && count < 1 {
			// fmt.Println(count)
			count++
		} else if val != num {
			num = val
			count = 0
		} else {
			nums[key] = math.MaxInt64
			returnableCount++
		}
	}
	sort.Ints(nums)
	return len(nums) - returnableCount
}
