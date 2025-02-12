package main

import (
	llist "cs/ds/linked"
	"fmt"
)

func main() {
	ll := llist.LinkedList{}

	var i uint8
	for i = 0; i < 10; i++ {
		ll.Append(int(i))
	}

	ll.Insert(-1, 0)
	ll.Insert(5, 10)

	num, err := ll.IndexOf(0)
	if err != nil {
		fmt.Printf("err => %v\n", err)
	} else {
		fmt.Printf("0 => %d\n", *num)
	}

	ll.Print()
}
