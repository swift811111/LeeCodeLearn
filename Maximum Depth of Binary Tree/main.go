package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Node struct {
	Dep   int
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	max := 0
	if root == nil {
		return 0
	}
	if root.Val == 0 && root.Right == nil && root.Left == nil {
		return 1
	}

	arNode := make([]Node, 0, 0)
	depth := 1

	for root.Left != nil || root.Right != nil || len(arNode) != 0 {
		if root.Right != nil {
			arNode = append(arNode, Node{depth + 1, root.Right})
		}
		if root.Left != nil {
			depth++
			if depth > max {
				max = depth
			}
			root = root.Left
			continue
		} else {
			depth = arNode[len(arNode)-1].Dep
			root = arNode[len(arNode)-1].Right
			arNode = arNode[:len(arNode)-1]
			if depth > max {
				max = depth
			}
		}

	}
	return max
}

// Input: root = [3,9,20,null,null,15,7]
// Output: 3
