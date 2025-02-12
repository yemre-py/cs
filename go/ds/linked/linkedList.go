package linked

import (
	"errors"
	"fmt"
)

/*
TODO: Refactor functions to accept all int type  paramaters as given, eliminate type conversions at main.go
*/

// create a new linked list
type LinkedList struct {
	head *Node
}

// create a new node
func createNode(val int) *Node {
	return &Node{
		Data: val,
		Next: nil,
	}
}

// print the list
func (ll *LinkedList) Print() {
	element := ll.head

	// traverse the list
	for element != nil {
		// if the next node is nil
		if element.Next == nil {
			fmt.Printf("%d", element.Data)
		} else {
			fmt.Printf("%d -> ", element.Data)
		}
		element = element.Next
	}
}

func (ll *LinkedList) Append(val int) error {
	newNode := createNode(val)

	// if list is empty
	if ll.head == nil {
		ll.head = newNode
		return nil
	}

	temp := ll.head

	// traverse to the last node
	for temp.Next != nil {
		temp = temp.Next
	}

	// append the new node
	temp.Next = newNode

	return nil
}

func (ll *LinkedList) Insert(val, ind int) error {
	newNode := createNode(val)

	// if index is invalid
	if ind < 0 {
		return errors.New("invalid index")
	}

	// if index is 0
	temp := ll.head
	if ind == 0 {
		newNode.Next = temp
		ll.head = newNode
		return nil
	}

	// traverse to the node before the insertion point
	index := 0
	for temp != nil && index < ind-1 {
		temp = temp.Next
		index++
	}

	// if index is out of bounds
	if temp == nil {
		return errors.New("index out of range")
	}

	// insert the new node
	newNode.Next = temp.Next
	temp.Next = newNode

	return nil
}

func (ll *LinkedList) IndexOf(ind int) (*int, error) {
	index := 0
	temp := ll.head

	for temp != nil && index < ind {
		temp = temp.Next
		index++
	}

	if temp == nil || ind < 0 {
		return nil, errors.New("index out of range")
	}

	return &temp.Data, nil
}
