package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if elem < root.Data {
			return BTreeSearchItem(root.Left, elem)
		} else if elem > root.Data {
			return BTreeSearchItem(root.Right, elem)
		} else {
			return root
		}
	} else {
		return nil
	}
}
