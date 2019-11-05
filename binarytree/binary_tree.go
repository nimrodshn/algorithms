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

func (b *BinaryTree) Min() *Node {
	node := b.root
	for node.left != nil {
		node = node.left
	}
	return node
}

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
func search(x *Node ,key int) *Node {
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
