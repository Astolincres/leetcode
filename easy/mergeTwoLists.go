package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	// Try to solve it here.
	var result ListNode
	var dir1, dir2, *ListNode
	dir1 = list1
	fmt.Println(dir1, dir1.Next.Val, dir2, result)
	list3 := &ListNode{
		Val: 1,
	}
	if list3.Next == nil {
		fmt.Println("Last Node")
	}
	for {
		break
	}	
	return nil
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

	fmt.Println(&list1.Next)
	fmt.Println(list1)
	fmt.Println(list2)
	// Your solution:
	result := mergeTwoList(list1, list2)
	fmt.Println(result)
}
