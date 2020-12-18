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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ar := make([]int, 0, 0)
	for l1 != nil || l2 != nil {
		if l1 == nil {
			ar = append(ar, l2.Val)
			l2 = l2.Next
			continue
		}
		if l2 == nil {
			ar = append(ar, l1.Val)
			l1 = l1.Next
			continue
		}
		if l1.Val >= l2.Val {
			ar = append(ar, l2.Val)
			l2 = l2.Next
		} else {
			ar = append(ar, l1.Val)
			l1 = l1.Next
		}
	}
	var h = &ListNode{}
	y := h
	for i := 0; i < len(ar); i++ {
		h.Next = &ListNode{}
		h = h.Next
		h.Val = ar[i]
	}

	return y.Next
}

// Input: l1 = [1,2,4], l2 = [1,3,4]
// Output: [1,1,2,3,4,4]
