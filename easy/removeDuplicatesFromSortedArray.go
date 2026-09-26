package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8, 8, 8, 9}
	fmt.Println(removeDuplicates(nums), nums)
}

func removeDuplicates(nums []int) int {
	var value int
	k := len(nums)
	if k <= 1 {
		return k
	}
	for i := 0; i < k; i++ {
		if nums[i] == nums[i+1] {
			value = nums[i+1]
			nums = append(nums[:i+1], nums[i+2:]...)
			nums = append(nums, value)
			k--
			i--
		}
	}

	return k
}
