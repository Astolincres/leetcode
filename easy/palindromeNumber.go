package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(12 % 100)
	fmt.Println(isPalindrome(123454321))
}

func isPalindrome(x int) bool {
	var newNumber int
	n := x
	if x < 0 {
		return false
	}
	if x/10 < 0 {
		return true
	}
	temp := []int{x % 10}
	for n/10 > 0 {
		n = n / 10
		temp = append(temp, n%10)
	}
	for i := 0; i < len(temp); i++ {
		newNumber = newNumber + temp[i]*int(math.Pow(10, float64(len(temp)-i-1)))
	}
	return newNumber-x == 0
}
