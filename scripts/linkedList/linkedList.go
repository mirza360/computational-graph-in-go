package main

import (
	"fmt"
)

type node struct {
	idx       int
	varA      float64
	varB      float64
	resVal    float64
	prev      *node
	next      *node
	childHead *bottomNode
	childTail *bottomNode
}

type list struct {
	tail *node
	head *node
}

type bottomNode struct {
	diffVal  float64
	upNode   *bottomNode
	downNode *bottomNode
}

func (l *list) Display() {
	//testDis()
	list := l.head
	for list != nil {
		fmt.Printf("value = %v and prev = %v and next= %v\n", list.idx, list.prev, list.next)
		list = list.next
	}
	fmt.Println()
}
func (n *node) testDis() {

	node := n.childHead
	for node != nil {
		fmt.Printf("value = %v and prev = %v and next= %v\n", node.diffVal, node.upNode, node.downNode)
		node = node.downNode
	}
	fmt.Println()
}

func (l *list) DisplayReverse() {
	list := l.tail
	for list != nil {
		fmt.Printf("value = %v\n", list.idx)
		list = list.prev
	}
	fmt.Println()
}

func (l *list) insertNode(newNode *node) {
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		l.tail = newNode
	}
}

func (n *node) addBottomNode(newNode *bottomNode) {
	if n.prev == nil {
		println("No parent Node")
	} else if n.childHead == nil {
		n.childHead = newNode
		n.childTail = newNode
	} else {
		currentNode := n.childHead
		for currentNode.downNode != nil {
			currentNode = currentNode.downNode
		}
		newNode.upNode = currentNode
		currentNode.downNode = newNode
		n.childTail = newNode
	}
}

func main() {
	/*items := &list{}
	size := 10
	//rand_number := make([]int, size, size)
	for i := 0; i < size; i++ {
		node := node{num: rand.Intn(100)}
		if node.num == 0 {
			i = i - 1
			continue

		}
		items.insertNode(&node)

		fmt.Printf("%v and number is %v\n", i, node.num)
	}
	items.Display()
	items := &list{}

	node1 := node{num: 1}
	items.insertNode(&node1)
	node2 := node{num: 2}
	items.insertNode(&node2)
	node3 := node{num: 3}
	items.addBottomNode(&node3)
	node4 := node{num: 4}
	items.addBottomNode(&node4)

	items.Display()*/
}
