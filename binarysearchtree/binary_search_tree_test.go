package binarysearchtree

import (
	"testing"
)

var (
	tree = New(&Node{
		key: 7,
	})
)

func TestInsert(t *testing.T) {
	tests := []struct {
		testCase    func()
		expectedErr error
	}{
		{
			testCase: func() {
				tree.Insert(&Node{
					key: 8,
				})
				if tree.root.right.key != 8 {
					t.Fatalf("Expected %d, instead recieved %d", 8, tree.root.right.key)
				}
				if tree.root.right.predecessor != tree.root {
					t.Fatalf("Expected %v, instead recieved %v", tree.root, tree.root.right.predecessor)
				}
			},
		},
		{
			testCase: func() {
				tree.Insert(&Node{
					key: 3,
				})
				if tree.root.left.key != 3 {
					t.Fatalf("Expected %d, instead recieved %d", 3, tree.root.left.key)
				}
				if tree.root.right.predecessor != tree.root {
					t.Fatalf("Expected %v, instead recieved %v", tree.root, tree.root.left.predecessor)
				}
			},
		},
		{
			testCase: func() {
				tree.Insert(&Node{
					key: 5,
				})
				if tree.root.left.right.key != 5 {
					t.Fatalf("Expected %d, instead recieved %d", 5, tree.root.left.right.key)
				}
				if tree.root.left.right.predecessor != tree.root.left {
					t.Fatalf("Expected %v, instead recieved %v", tree.root, tree.root.left.right.predecessor)
				}
			},
		},
		{
			testCase: func() {
				tree.Insert(&Node{
					key: 2,
				})
				if tree.root.left.left.key != 2 {
					t.Fatalf("Expected %d, instead recieved %d", 2, tree.root.left.left.key)
				}
				if tree.root.left.left.predecessor != tree.root.left {
					t.Fatalf("Expected %v, instead recieved %v", tree.root, tree.root.left.right.predecessor)
				}
			},
		},
		{
			testCase: func() {
				tree.Insert(&Node{
					key: 6,
				})
				if tree.root.left.right.right.key != 6 {
					t.Fatalf("Expected %d, instead recieved %d", 6, tree.root.left.right.right.key)
				}
				if tree.root.left.right.right.predecessor != tree.root.left.right {
					t.Fatalf("Expected %v, instead recieved %v", tree.root.left.right, tree.root.left.right.predecessor)
				}
			},
		},
		{
			testCase: func() {
				tree.Insert(&Node{
					key: 6,
				})
				if tree.root.left.right.right.key != 6 {
					t.Fatalf("Expected %d, instead recieved %d", 6, tree.root.left.right.right.key)
				}
				if tree.root.left.right.right.predecessor != tree.root.left.right {
					t.Fatalf("Expected %v, instead recieved %v", tree.root.left.right, tree.root.left.right.predecessor)
				}
			},
		},
	}
	for _, test := range tests {
		test.testCase()
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		testCase    func()
		expectedErr error
	}{
		{
			testCase: func() {
				_, err := tree.Search(8)
				if err != nil {
					t.Fatalf("An error occurred while trying to search of %d:%v", 8, err)
				}
			},
		},
		{
			testCase: func() {
				_, err := tree.Search(6)
				if err != nil {
					t.Fatalf("An error occurred while trying to search of %d:%v", 8, err)
				}
			},
		},
		{
			testCase: func() {
				_, err := tree.Search(1)
				if err != nil && err != ErrNotFound {
					t.Fatalf("An unexpected error occurred while trying to search of %d:%v", 1, err)
				}
			},
		},
	}
	for _, test := range tests {
		test.testCase()
	}
}

func TestMax(t *testing.T) {
	node := tree.Max()
	if node.key != 8 {
		t.Fatalf("Expected %d as maximum, instead found: %d", 8, node.key)
	}
}

func TestMin(t *testing.T) {
	node := tree.Min()
	if node.key != 2 {
		t.Fatalf("Expected %d as minimum, instead found: %d", 2, node.key)
	}
}

func TestDelete(t *testing.T) {
	tree.Delete(tree.root.left)
	if tree.root.left.key != 5 {
		t.Fatalf("Expected %d as roots left node, instead found: %d", 5, tree.root.left.key)
	}
	if tree.root.left.right.key != 6 {
		t.Fatalf("Expected %d , instead found: %d", 6, tree.root.left.right.key)
	}
	if tree.root.left.right.predecessor.key != 5 {
		t.Fatalf("Expected %d , instead found: %d", 5, tree.root.left.right.predecessor.key)
	}
}
