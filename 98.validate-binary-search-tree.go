package main

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	current := root
	var prev *int

	// using inorder traversal technique from problem 94.binary-tree-inorder-traversal
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil && current.Val <= *prev {
			return false
		}

		prev = &current.Val
		current = current.Right
	}

	return true

	// other approach (cleaner)
	// return validate(root, math.MinInt, math.MaxInt)
}

func validate(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max && root.Val <= min {
		return false
	}

	return validate(root.Left, min, root.Val) && validate(root.Right, root.Val, max)
}
