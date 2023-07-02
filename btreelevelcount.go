package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root != nil {
		leftHeight := BTreeLevelCount(root.Left)
		rightHeight := BTreeLevelCount(root.Right)

		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	} else {
		return 0
	}
}
