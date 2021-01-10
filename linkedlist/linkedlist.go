package linkedlist

import (
	"errors"
	"fmt"
)

var errNoData = errors.New("No data in the Linkedlist")

// Node struct
type Node struct {
	data int
	next *Node
}

// LinkedList struct
type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Length of the linkedlist
func (l LinkedList) Length() int {
	return l.length
}

// Add a data into the linkedlist
func (l *LinkedList) Add(data int) {
	if l.length == 0 {
		l.head = &Node{data: data, next: nil}
		l.tail = l.head
	} else {
		newNode := &Node{data: data, next: nil}
		l.tail.next = newNode
		l.tail = newNode
	}
	l.length++
}

// Remove a data in the linkedlist - first meet
func (l *LinkedList) Remove(data int) error {
	if l.length == 0 {
		return errNoData
	}

	cursor := l.head
	if cursor.data == data {
		targetNode := cursor
		l.head = l.head.next
		fmt.Println(targetNode)
		targetNode = nil
		l.length--
		return nil
	}

	// not head
	for cursor.next != nil {
		if cursor.next.data == data {
			targetNode := cursor.next
			if cursor.next.next == nil {
				// tail
				cursor.next = nil
				l.tail = cursor
			} else {
				cursor.next = cursor.next.next
			}
			fmt.Println(targetNode)
			targetNode = nil
			l.length--
			break
		}
		cursor = cursor.next
	}
	return nil
}

// Traverse the linkedlist
func (l LinkedList) Traverse() {
	cursor := l.head
	for cursor != nil {
		fmt.Println(cursor.data)
		cursor = cursor.next
	}
}
