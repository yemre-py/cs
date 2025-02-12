package main

import (
	llist "cs/ds/linked"
)

func main() {
	ll := llist.LinkedList{}

	var i uint8
	for i=0;i<10;i++{
		ll.Append(int(i))
	}

	ll.Print()
}
