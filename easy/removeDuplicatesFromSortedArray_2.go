package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8, 8, 8, 9}
	fmt.Println(removeDuplicates(nums), nums)
}

func removeDuplicates(nums []int) int {
	k := 0

	newArray, duplicateArray := []int{}, []int{}

	currentValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentValue == nums[i] {
			duplicateArray = append(duplicateArray, currentValue)
		} else {
			newArray = append(newArray, currentValue)
			currentValue = nums[i]
			k++
		}
	}
	nums = append(newArray, duplicateArray...)
	return k
}
