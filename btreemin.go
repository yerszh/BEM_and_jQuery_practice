package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root != nil {
		node := root
		for node.Left != nil {
			node = node.Left
		}
		return node
	} else {
		return nil
	}
}
