package main

type AVL_Node struct {
	value     int
	RightNode *AVL_Node
	LeftNode  *AVL_Node
	heigh     int
}

func Insert(root *AVL_Node, value int) *AVL_Node {
	if root == nil {
		return &AVL_Node{value: value, RightNode: nil, LeftNode: nil, heigh: 0}
	}
	if value > root.value {
		root.RightNode = Insert(root.RightNode, value)
	} else {
		root.LeftNode = Insert(root.LeftNode, value)
	}
	root.heigh = max(height(root.LeftNode), height(root.RightNode)) + 1

	return balance(root)
}

func balance(root *AVL_Node) *AVL_Node {
	if isRightHeavy(root) {
		if balanceFactor(root.RightNode) > 0 {
			root.RightNode = rightRotation(root.RightNode)
			return leftRotation(root)
		} else if balanceFactor(root.RightNode) < 0 {
			return leftRotation(root)
		}
	} else if isLeftHeavy(root) {
		if balanceFactor(root.LeftNode) > 0 {
			return rightRotation(root)
		} else if balanceFactor(root.LeftNode) < 0 {
			root.LeftNode = leftRotation(root.LeftNode)
			return rightRotation(root)
		}
	}
	return root

}
func leftRotation(root *AVL_Node) *AVL_Node {
	right := root.RightNode
	root.RightNode = right.LeftNode
	right.LeftNode = root

	root.heigh = max(height(right.LeftNode), height(right.RightNode)) + 1
	root.heigh = max(height(root.LeftNode), height(root.RightNode)) + 1

	return right
}

func rightRotation(root *AVL_Node) *AVL_Node {
	newRoot := root.LeftNode
	root.LeftNode = newRoot.RightNode
	newRoot.RightNode = root

	root.heigh = max(height(newRoot.LeftNode), height(newRoot.RightNode)) + 1
	root.heigh = max(height(root.LeftNode), height(root.RightNode)) + 1

	return newRoot
}
func isLeftHeavy(root *AVL_Node) bool {
	return balanceFactor(root) >= 1

}

func isRightHeavy(root *AVL_Node) bool {
	return balanceFactor(root) <= -1
}
func balanceFactor(node *AVL_Node) int {
	if node == nil {
		return 0
	} else {
		return height(node.LeftNode) - height(node.RightNode)
	}
}
func height(node *AVL_Node) int {
	if node == nil {
		return -1
	} else {
		return node.heigh
	}
}
func max(a int, b int) int {

	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	tree := AVL_Node{value: 10, RightNode: nil, LeftNode: nil, heigh: 0}
	Insert(&tree, 5)
	Insert(&tree, 6)

}
