package binarysearchtree

import "errors"

var (
	// ErrNotFound represent a Node not found error.
	ErrNotFound = errors.New("Node not found")
)

// BinarySearchTree is a representation of a binary tree
type BinarySearchTree struct {
	root *Node
}

// Node represents a node in the binary tree.
type Node struct {
	key         int
	predecessor *Node
	left        *Node
	right       *Node
}

// New returns a new BinarySearchTree rooted with root.
func New(root *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: root,
	}
}

// Min returns the node with the minimum key.
func (b *BinarySearchTree) Min() *Node {
	node := b.root
	for node.left != nil {
		node = node.left
	}
	return node
}

// Max returns the node with the maximum key.
func (b *BinarySearchTree) Max() *Node {
	node := b.root
	for node.right != nil {
		node = node.right
	}
	return node
}

// Search returns the node in the Binary Tree with a given key.
func (b *BinarySearchTree) Search(key int) (*Node, error) {
	return search(b.root, key)
}

// A helper function for Search.
func search(x *Node, key int) (*Node, error) {
	for x != nil {
		if x.key == key {
			return x, nil
		}
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return nil, ErrNotFound
}

// Insert inserts a node into the binary tree.
func (b *BinarySearchTree) Insert(node *Node) {
	var pred *Node
	x := b.root
	for x != nil {
		pred = x
		if node.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	node.predecessor = pred
	if pred == nil {
		b.root = pred
	}
	if node.key < pred.key {
		pred.left = node
	} else {
		pred.right = node
	}
}

// Successor returns the successor of a given node (with respect to its key).
func (b *BinarySearchTree) Successor(node *Node) *Node {
	if node.right != nil {
		rightSubTree := BinarySearchTree{node.right}
		return rightSubTree.Min()
	}
	y := node.predecessor
	for y != nil && node == y.right {
		node = y
		y = y.predecessor
	}
	return y
}

// Delete delets a node from the binary tree.
func (b *BinarySearchTree) Delete(node *Node) {
	var y *Node
	var x *Node
	// If either of the nodes children is nil.
	// We simply remove the node and splice in it's children.
	// If both are non-nil we replace node with its successor.
	if node.right == nil || node.left == nil {
		y = node
	} else {
		y = b.Successor(node)
	}
	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}
	// Check end-case of y being the root.
	if y.predecessor == nil {
		b.root = x
	} else {
		if y == y.predecessor.left {
			y.predecessor.left = x
		} else {
			y.predecessor.right = x
		}
	}
	// Replace node's key with its successors key.
	if y != node {
		node.key = y.key
	}
}
