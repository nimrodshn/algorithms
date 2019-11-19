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
