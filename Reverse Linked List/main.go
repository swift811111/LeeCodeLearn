package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ar := make([]int, 0, 0)
	for head != nil {
		ar = append(ar, head.Val)
		head = head.Next
	}
	var h = &ListNode{}
	y := h
	for i := len(ar) - 1; i >= 0; i-- {
		h.Next = &ListNode{}
		h = h.Next
		h.Val = ar[i]
	}
	return y.Next
}

// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
