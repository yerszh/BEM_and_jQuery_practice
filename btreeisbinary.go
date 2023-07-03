package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root != nil {
		switch {
		case root.Left != nil && root.Left.Data > root.Data:
			return false

		case root.Right != nil && root.Right.Data < root.Data:
			return false
		case !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right):
			return false

		}
	}
	return true
}
