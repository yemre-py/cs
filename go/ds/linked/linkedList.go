package linked

import "fmt"

/*
	TODO: Refactor functions to accept all int type  paramaters as given, eliminate type conversions at main.go
	

*/
type LinkedList struct {
	head *Node
}

func createNode(val int) *Node {
	return &Node{
		Data: val,
	}
}

func (ll *LinkedList) Print() {
	element := ll.head
	for element != nil {
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

	if ll.head == nil {
		ll.head = newNode
		return nil
	}

	temp := ll.head

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = newNode

	return nil
}

func (ll *LinkedList) Insert(val, ind int) error {

}

