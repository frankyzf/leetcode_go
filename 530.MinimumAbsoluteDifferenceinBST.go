package main

func getMinimumDifference(root *TreeNode) int {
	if root != nil {
		return getMinForNonNullTree(root)
	}
	return 100000
}

func getMinForNonNullTree(root *TreeNode) int {
	minDiff := 100000
	if root.Left != nil {
		parentNode := root.Left //get the rightest node
		for parentNode.Right != nil {
			parentNode = parentNode.Right
		}
		if root.Val-parentNode.Val < minDiff {
			minDiff = root.Val - parentNode.Val
		}
		leftMin := getMinForNonNullTree(root.Left)
		if leftMin < minDiff {
			minDiff = leftMin
		}
	}
	if root.Right != nil {
		parentNode := root.Right
		for parentNode.Left != nil {
			parentNode = parentNode.Left
		}
		if parentNode.Val-root.Val < minDiff {
			minDiff = parentNode.Val - root.Val
		}
		rightMin := getMinForNonNullTree(root.Right)
		if rightMin < minDiff {
			minDiff = rightMin
		}
	}
	return minDiff
}
