package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hashPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	newSum := targetSum - root.Val

	return hashPathSum(root.Left, newSum) || hashPathSum(root.Right, newSum)
}
