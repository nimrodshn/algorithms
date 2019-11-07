package binarytree

// BinaryTree is a representation of a binary tree
type BinaryTree struct {
	root *Node
}

// Node represents a node in the binary tree.
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Min returns the node with the minimum key.
func (b *BinaryTree) Min() *Node {
	node := b.root
	for node.left != nil {
		node = node.left
	}
	return node
}

// Max returns the node with the maximum key.
func (b *BinaryTree) Max() *Node {
	node := b.root
	for node.right != nil {
		node = node.right
	}
	return node
}

// Search returns the node in the Binary Tree with a given key.
func (b *BinaryTree) Search(key int) *Node {
	return search(b.root, key)
}

// A helper function for Search.
func search(x *Node, key int) *Node {
	for x != nil {
		if x.key == key {
			return x
		}
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return nil
}

// Insert inserts a node into the binary tree.
func (b *BinaryTree) Insert(node *Node) {
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
	if pred == nil {
		b.root = pred
	}
	if node.key < pred.key {
		pred.left = node
	} else {
		pred.right = node
	}
}
