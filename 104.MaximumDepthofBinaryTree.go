package main

func maxDepth(root *TreeNode) int {
	return md(root)
}

func md(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	if root.Left != nil {
		depth = md(root.Left)
	}
	if root.Right != nil {
		d2 := md(root.Right)
		if d2 > depth {
			depth = d2
		}
	}
	return depth + 1
}
