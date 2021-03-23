package main

func main() {

	NodeDemoWithMethods()
	NodeDemoWithRegularFunctions()

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
