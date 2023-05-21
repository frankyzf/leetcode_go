package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortTree(nums, 0, len(nums)-1)
}

func sortTree(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	} else if start == end {
		return &TreeNode{
			Val: nums[start],
		}
	} else {
		mid := int((start + end + 1) / 2)
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = sortTree(nums, start, mid-1)
		root.Right = sortTree(nums, mid+1, end)
		return root
	}
}
