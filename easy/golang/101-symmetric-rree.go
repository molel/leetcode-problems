package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricSubtree(root.Left, root.Right)
}

func isSymmetricSubtree(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	} else {
		return left.Val == right.Val && isSymmetricSubtree(left.Left, right.Right) && isSymmetricSubtree(left.Right, right.Left)
	}
}
