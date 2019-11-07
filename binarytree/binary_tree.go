package binarytree

// BinaryTree is a representation of a binary tree
type BinaryTree struct {
	root *Node
}

// Node represents a node in the binary tree.
type Node struct {
	key         int
	predecessor *Node
	left        *Node
	right       *Node
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
func (b *BinaryTree) Successor(node *Node) *Node {
	if node.right != nil {
		rightSubTree := BinaryTree{node.right}
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
func (b *BinaryTree) Delete(node *Node) {
	y := node.predecessor
	// Both of the nodes children are empty (i.e. node is a leaf).
	// We simply remove the node.
	if node.right == nil && node.left == nil {
		if y.right == node {
			y.right = nil
		} else {
			y.left = nil
		}
		return
	}
	// One of the nodes children is nil.
	// Splice out the node and attach its non-nil child to his parent.
	if node.right == nil {
		if y.right == node {
			y.right = node.left
		} else {
			y.left = node.left
		}
		return
	} else if node.left == nil {
		if y.right == node {
			y.right = node.right
		} else {
			y.left = node.right
		}
		return
	}
	// Both of the nodes are non-nil.
	successor := b.Successor(node)
	var x *Node
	// fix childrens predecessor.
	if y.left != nil {
		x = successor.left
	} else {
		x = successor.right
	}
	x.predecessor = successor.predecessor
	if successor != node {
		// replace the node with its successor.
		node.key = successor.key
	}
}
