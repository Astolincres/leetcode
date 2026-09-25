package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	// Try to solve it here.
	var dir1, dir2, result, tail *ListNode
	dir1 = list1
	dir2 = list2
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if dir1.Val < dir2.Val {
		result = &ListNode{Val: dir1.Val}
		dir1 = dir1.Next
	} else {
		result = &ListNode{Val: dir2.Val}
		dir2 = dir2.Next
	}
	tail = result
	for {
		if dir1.Next == nil && dir2.Next == nil {
			return result
		}
		if dir1.Next == nil && dir2.Next != nil {
			for {
				tail.Next = dir2
				tail = tail.Next
				if dir2.Next != nil {
					dir2 = dir2.Next
				} else {
					return result
				}
			}
		}
		if dir1.Next != nil && dir2.Next == nil {
			for {
				tail.Next = dir1
				tail = tail.Next
				if dir1.Next != nil {
					dir1 = dir1.Next
				} else {
					return result
				}
			}
		}
		if dir1.Val > dir2.Val {
			tail.Next = dir2
			tail = tail.Next
			dir2 = dir2.Next
		} else {
			tail.Next = dir1
			tail = tail.Next
			dir1 = dir1.Next
		}
	}
}

func main() {
	// list1: 1 -> 3 -> 5 -> 8
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 8,
				},
			},
		},
	}

	// list2: 2 -> 4 -> 6 -> 7
	list2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 7,
				},
			},
		},
	}

	// Your solution:
	result := mergeTwoList(list1, list2)
	fmt.Println(result)
	for result.Next != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
