package redblacktrees

// Color is the color of a node in the tree.
type Color string

const (
	red   Color = "red"
	black Color = "black"
)

// RedBlackTree is an implementation of the red-black tree data structure.
type RedBlackTree struct {
	root *Node
}

// Node represents a node in the tree
type Node struct {
	left   *Node
	right  *Node
	color  Color
	parent *Node
}

func (t *RedBlackTree) RotateLeft(node *Node) {
	y := node.right
	node.right = y.left
	if y.left != nil {
		y.left.parent = node
	}
	y.parent = node.parent
	node.parent = y
	if t.root == node {
		t.root = y
	} else {
		if node.parent.right == node {
			node.parent.right = y
		} else {
			node.parent.left = y
		}
	}
}

func (t *RedBlackTree) RotateRight(node *Node) {
	y := node.left
	node.left = y.right
	if y.right != nil {
		y.right.parent = node
	}
	y.parent = node.parent
	node.parent = y
	if t.root == node {
		t.root = y
	} else {
		if node.parent.right == node {
			node.parent.right = y
		} else {
			node.parent.left = y
		}
	}
}