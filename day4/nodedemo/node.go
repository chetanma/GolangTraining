package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

//initializer
func NewNode(val int) Node {
	return Node{
		Value: val,
		Next:  nil,
	}
}

//Methods: AddElement & TraverseAndDisplay

func (n *Node) AddElement(ele int) {

	temp := n

	for temp.Next != nil {
		temp = temp.Next
	}
	newNode := NewNode(ele)
	temp.Next = &newNode
}

func (n *Node) TraverseAndDisplay() {

	temp := n

	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}

}

func (n *Node) TraverseAsync() <-chan int {
	//A: create a channel
	ch := make(chan int)

	//B: Just invoke concurrent traversal func.
	go func() {
		//D: Traverse the list and send values over channel
		defer close(ch)
		temp := n
		for temp != nil {
			ch <- temp.Value
			temp = temp.Next
		}

	}()

	//C: return the channel to the caller
	return ch
}

// Regular functions:
func AddElementFn(n *Node, ele int) {
	temp := n

	for temp.Next != nil {
		temp = temp.Next
	}
	newNode := NewNode(ele)
	temp.Next = &newNode
}

func TraverseAndDisplayFn(n *Node) {

	temp := n

	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}

}
