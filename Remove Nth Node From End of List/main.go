package main

import "fmt"

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ar := make([]int, 0, 4)
	for head.Next != nil {
		ar = append(ar, head.Val)
		head = head.Next
	}
	ar = append(ar, head.Val)
	var h = &ListNode{}
	y := h
	if len(ar) == 1 && n == 1 {
		return nil
	}
	for i := 0; i < len(ar); i++ {
		if i == len(ar)-n {
			continue
		}
		fmt.Println(i)
		h.Next = &ListNode{}
		h = h.Next
		h.Val = ar[i]
	}
	return y.Next
}

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
