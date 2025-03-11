package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	temp := this.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	n := &Node{
		val:  val,
		next: nil,
	}
	if this.size == 0 {
		this.head = n
		this.tail = n
		this.size++
		return
	}
	n.next = this.head
	this.head = n
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := &Node{
		val:  val,
		next: nil,
	}
	if this.size == 0 {
		this.head = n
		this.tail = n
		this.size++
		return
	}
	this.tail.next = n
	this.tail = n
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	} else if index == this.size {
		this.AddAtTail(val)
		return
	} else if index > this.size || index < 0 {
		return
	}
	temp := this.head
	for i := 0; i < index-1; i++ {
		temp = temp.next
	}
	n := &Node{
		val:  val,
		next: temp.next,
	}
	temp.next = n
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	} else if index == 0 {
		this.head = this.head.next
		this.size--
		return
	}
	temp := this.head
	for i := 0; i < index-1; i++ {
		if temp.next == nil {
			return
		}
		temp = temp.next
	}
	temp.next = temp.next.next
	if index == this.size-1 {
		this.tail = temp
	}
	this.size--
}

func (this *MyLinkedList) print() {
	for n := this.head; n != this.tail; n = n.next {
		fmt.Printf("%d -> ", n.val)
	}
	fmt.Printf("%d\n", this.tail.val)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.print()
	obj.AddAtTail(3)
	obj.print()
	obj.AddAtIndex(1, 2)
	obj.print()
	fmt.Println(obj.Get(1))
	obj.DeleteAtIndex(2)
	obj.print()
	fmt.Println(obj.Get(0))
}
