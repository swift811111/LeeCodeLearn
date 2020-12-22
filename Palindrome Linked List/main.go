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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	ar := make([]int, 0, 0)
	for head != nil {
		ar = append(ar, head.Val)
		head = head.Next
	}
	n := len(ar) / 2
	for i := 0; i < n; i++ {
		if ar[i] != ar[len(ar)-1-i] {
			return false
		}
	}
	return true
}

// Input: 1->2->2->1
// Output: true
