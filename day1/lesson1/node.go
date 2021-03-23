package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func Node_init() {
	//initialization
	node0 := Node{
		Value: 20,
		Next:  nil,
	}
	node := Node{
		Value: 10,
		Next:  &node0,
	}

	fmt.Printf("%v %v", node, *node.Next)
}
