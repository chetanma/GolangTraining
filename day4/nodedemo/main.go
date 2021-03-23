package main

import "fmt"

func main() {

	//NodeDemoWithMethods()
	//NodeDemoWithRegularFunctions()
	NodeDemoWithChannel()

}

func NodeDemoWithChannel() {
	rNode1 := NewNode(100)
	rNode1.AddElement(200)
	rNode1.AddElement(300)
	rNode1.AddElement(55)

	outCh := rNode1.TraverseAsync()

	//can lanuch lot of workers to process the data
	for out := range outCh {
		fmt.Println(out)
	}
}

func NodeDemoWithMethods() {
	rNode1 := NewNode(100)
	rNode1.AddElement(200)
	rNode1.AddElement(300)
	rNode1.AddElement(55)
	rNode1.TraverseAndDisplay()
}

func NodeDemoWithRegularFunctions() {
	rNode2 := NewNode(100)
	AddElementFn(&rNode2, 53)
	AddElementFn(&rNode2, 54)
	AddElementFn(&rNode2, 64)
	TraverseAndDisplayFn(&rNode2)
}
