package mysort

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(t *TreeNode) {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	PrintTree(t.Left)
	PrintTree(t.Right)
}

func PrintTreeByStack(t *TreeNode) {

}

func PrintTreeI(t *TreeNode) {

}
