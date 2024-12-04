package main

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func searchRangeBst(root *TreeNode, low, high int) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	if root.val >= low && root.val <= high {
		result = append(result, root.val)
	}

	if root.val > low {
		result = append(result, searchRangeBst(root.left, low, high)...)
	}

	if root.val < high {
		result = append(result, searchRangeBst(root.right, low, high)...)
	}

	return result
}

func main() {
	root := &TreeNode{val: 5}
	root.left = &TreeNode{val: 3}
	root.left.left = &TreeNode{val: 1}
	root.left.right = &TreeNode{val: 4}
	root.right = &TreeNode{val: 8}
	root.right.right = &TreeNode{val: 10}

	low := 4
	high := 9
	result := searchRangeBst(root, low, high)
	slices.SortFunc(result, CmpDesc)
	fmt.Println(result)
}

func CmpDesc(x, y int) int {
	if x > y {
		return -1
	} else if x < y {
		return 1
	} else {
		return 0
	}
}
