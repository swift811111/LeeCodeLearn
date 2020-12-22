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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	ar := make(map[*ListNode]bool)
	for head != nil {
		if ar[head.Next] {
			return true
		}
		ar[head.Next] = true
		head = head.Next
	}
	return false
}

// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.
