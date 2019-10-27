package linkedlist

import (
	"errors"
)

var (
	// ErrNotFound represent a link not found error.
	ErrNotFound = errors.New("Link not found")
)

// Link is a link implementation.
type Link struct {
	key  int
	next *Link
}

// LinkedList is a linked list implementation.
type LinkedList struct {
	head *Link
}

// New returns a new linked list.
func New() *LinkedList {
	return new(LinkedList)
}

// Add adds a link with key k to the head of the list.
func (l *LinkedList) Add(k int) {
	link := &Link{
		key:  k,
		next: l.head,
	}
	l.head = link
}

// Remove removes the first link with a given key from the head of the list.
func (l *LinkedList) Remove(key int) error {
	curr := l.head
	if curr.key == key {
		l.head = l.head.next
		return nil
	}
	for curr.next != nil {
		if curr.next.key == key {
			curr.next = curr.next.next
			return nil
		}
		curr = curr.next
	}
	return ErrNotFound
}

// Search searches for the first link with a given key in the list.
func (l *LinkedList) Search(key int) (*Link, error) {
	curr := l.head
	for curr != nil {
		if curr.key == key {
			return curr, nil
		}
		curr = curr.next
	}
	return nil, ErrNotFound
}
