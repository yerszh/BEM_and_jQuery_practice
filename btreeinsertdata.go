package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	} else {
		switch {
		case data > root.Data:
			root.Right = BTreeInsertData(root.Right, data)
			root.Right.Parent = root
		case data < root.Data:
			root.Left = BTreeInsertData(root.Left, data)
			root.Left.Parent = root
		case data == root.Data:
			root.Parent = BTreeInsertData(root.Parent, data)
		}
		return root
	}
}
