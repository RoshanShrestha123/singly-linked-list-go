// I am just trying to understand the linkedList on the golang,
// please don't judge my code 🙃

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) addToFirst(data int) {
	newData := Node{data: data, next: nil}

	if l.head != nil {
		newData.next = l.head
		l.head = &newData
	} else {
		l.head = &newData
	}
}

func (l *LinkedList) deleteValue(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		if currentNode.next.data == value {
			currentNode.next = currentNode.next.next
			return
		} else {
			currentNode = currentNode.next
		}
	}

	if currentNode.next == nil {
		fmt.Println("Data not found to delete!")
	}

}

func (l *LinkedList) addToLast(data int) {
	node := Node{data: data, next: nil}

	if l.head == nil {
		l.head = &node
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = &node
}

func (l *LinkedList) displayOutput() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	fmt.Println("LinkedList implementation")
	list := LinkedList{}
	list.addToLast(2000)
	list.addToFirst(5000)
	list.addToLast(6000)
	list.addToFirst(7000)
	list.addToFirst(8000)
	list.deleteValue(2000)
	list.displayOutput()

}
