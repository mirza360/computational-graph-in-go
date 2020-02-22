package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	num  int
	prev *node
	next *node
}

type list struct {
	tail  *node
	start *node
}

func (l *list) Display() {
	list := l.start
	for list != nil {
		fmt.Printf("value = %v and prev = %v and next= %v\n", list.num, list.prev, list.next)
		list = list.next
	}
	fmt.Println()
}

func (l *list) DisplayReverse() {
	list := l.tail
	for list != nil {
		fmt.Printf("value = %v\n", list.num)
		list = list.prev
	}
	fmt.Println()
}

func (l *list) insertNode(newNode *node) {
	if l.start == nil {
		l.start = newNode
		l.tail = newNode
	} else {
		currentNode := l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		l.tail = newNode
	}
}

func main() {
	items := &list{}
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
}
