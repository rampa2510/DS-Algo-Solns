package main

import "fmt"

type node struct {
	data    int
	pointer *node
}

type linkedList struct {
	head *node
}

func createNewNode(data int) *node {
	newNode := &node{data: data, pointer: nil}
	// fmt.Println("create")
	// fmt.Println(newNode)
	return newNode
}

func (ll *linkedList) isEmpty() bool {
	if ll.head == nil {
		return true
	}
	return false
}

func (ll *linkedList) insertAtEnd(data int) {
	var newNode *node
	newNode = new(node)
	newNode = createNewNode(data)
	// fmt.Println("ins")
	// fmt.Println(newNode)
	if ll.head == nil {
		ll.head = newNode
		return
	}
	endNode := ll.head
	for endNode.pointer != nil {
		endNode = endNode.pointer
	}
	endNode.pointer = newNode

}

func (ll *linkedList) insertAtBeg(data int) {
	newNode := createNewNode(data)
	newNode.pointer = ll.head
	ll.head = newNode
}

func (ll *linkedList) insert(data, pos int) {
	newNode := createNewNode(data)
	if ll.isEmpty() && pos > 1 {
		fmt.Println("Empty linked list")
		return
	}
	var currNode *node
	currNode = ll.head
	for i := 1; i < pos-1; i++ {
		currNode = currNode.pointer
	}
	newNode.pointer = currNode.pointer
	currNode.pointer = newNode
}

func (ll *linkedList) display() {
	if ll.isEmpty() {
		fmt.Println("Linked list empty")
		return
	}
	currNode := ll.head
	for currNode != nil {
		fmt.Println(currNode.data)
		currNode = currNode.pointer
	}
}

func (ll *linkedList) deleteAtBeg() *node {
	if ll.isEmpty() {
		fmt.Println("Empty linked list")
	}
	nodeToDelete := ll.head
	ll.head = nodeToDelete.pointer
	return nodeToDelete
}

func (ll *linkedList) delete(pos int) *node {
	if ll.isEmpty() {
		fmt.Println("Empty linked list")
		return nil
	}
	var currNode *node
	currNode = ll.head
	var nodeTODelete *node
	nodeTODelete = new(node)
	for i := 1; i < pos-1; i++ {
		if currNode.pointer == nil {
			fmt.Println("Position exceeds lined list length")
			return nil
		}
		currNode = currNode.pointer
	}
	nodeTODelete = currNode.pointer
	currNode.pointer = nodeTODelete.pointer
	return nodeTODelete
}

func (ll *linkedList) update(pos, data int) {
	if ll.isEmpty() {
		fmt.Println("Empty linked list")
		return
	}
	currNode := ll.head
	for i := 1; i < pos; i++ {
		currNode = currNode.pointer
	}
	currNode.data = data
}

func (ll *linkedList) reverseLL() {
	if ll.isEmpty() {
		fmt.Println("Empty linked list")
		return
	}
	currNode := ll.head.pointer
	revLL := ll.head
	revLL.pointer = nil
	// fmt.Println(currNode, revLL)
	var nextNode *node
	nextNode = new(node)
	for currNode != nil {
		nextNode = currNode.pointer // err
		currNode.pointer = revLL
		revLL = currNode
		currNode = nextNode
	}
	ll.head = revLL
	// fmt.Println(ll.head)
}

func main() {
	var ll *linkedList
	ll = new(linkedList)
	ll.display()
	ll.insertAtBeg(1)
	ll.insertAtEnd(2)
	ll.insertAtEnd(3)
	ll.insertAtEnd(4)
	ll.insertAtEnd(5)
	fmt.Println("displat full")
	ll.display()
	ll.delete(2)
	fmt.Println("display after delete")
	ll.display()
	ll.insert(6, 3)
	fmt.Println("display after inserting in b/w")
	ll.display()
	ll.deleteAtBeg()
	fmt.Println("Display after deleteing at beg")
	ll.display()
	ll.update(2, 11)
	fmt.Println("Display after updating")
	ll.display()
	ll.reverseLL()
	fmt.Println("Display after reverse")
	ll.display()
}
