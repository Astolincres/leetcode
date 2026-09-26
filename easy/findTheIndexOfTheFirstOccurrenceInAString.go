package main

import "fmt"

func main() {
	needle := "word"
	haystack := "some sentence that may include word"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	k := len(needle)
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if k == 1 {
				return i
			}
			for j := 1; j < k && i+j < len(haystack); j++ {
				if haystack[i+j] != needle[j] {
					break
				}
				if j == k-1 {
					return i
				}
			}
		}
	}
	return -1
}
