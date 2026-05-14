package list

import "errors"

type Node2P struct {
	prev  *Node2P
	value int
	next  *Node2P
}

type DoublyLinkedList struct {
	head     *Node2P
	tail     *Node2P
	inserted int
}

func (list *DoublyLinkedList) Add(value int) {
	newNode := &Node2P{value: value, prev: list.tail}

	if list.tail != nil {
		list.tail.next = newNode
	} else {
		list.head = newNode
	}

	list.tail = newNode
	list.inserted++
}

func (list *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo")
	}
	if index > list.inserted {
		return errors.New("index acima da faixa aceitável")
	}

	if index == list.inserted {
		list.Add(value)
		return nil
	}

	newNode := &Node2P{value: value}

	if index == 0 {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
		list.inserted++
		return nil
	}

	nextNode := list.head
	for i := 0; i < index; i++ {
		nextNode = nextNode.next
	}

	prevNode := nextNode.prev
	newNode.prev = prevNode
	newNode.next = nextNode
	prevNode.next = newNode
	nextNode.prev = newNode

	list.inserted++
	return nil
}

func (list *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo")
	}
	if index >= list.inserted {
		return errors.New("index acima da faixa aceitável")
	}

	curNode := list.head
	for i := 0; i < index; i++ {
		curNode = curNode.next
	}

	if curNode.prev != nil {
		curNode.prev.next = curNode.next
	} else {
		list.head = curNode.next
	}

	if curNode.next != nil {
		curNode.next.prev = curNode.prev
	} else {
		list.tail = curNode.prev
	}

	list.inserted--
	return nil
}

func (list *DoublyLinkedList) Get(index int) (int, error) {
	if index < 0 {
		return -1, errors.New("index não pode ser negativo")
	}
	if index >= list.inserted {
		return -1, errors.New("index acima da faixa aceitável")
	}

	curNode := list.head
	for i := 0; i < index; i++ {
		curNode = curNode.next
	}
	return curNode.value, nil
}

func (list *DoublyLinkedList) Set(value int, index int) error {
	if index < 0 {
		return errors.New("index não pode ser negativo")
	}
	if index >= list.inserted {
		return errors.New("index acima da faixa aceitável")
	}

	curNode := list.head
	for i := 0; i < index; i++ {
		curNode = curNode.next
	}
	curNode.value = value
	return nil
}

func (list *DoublyLinkedList) Size() int {
	return list.inserted
}

func (list *DoublyLinkedList) Reverse() {
	for curNode := list.head; curNode != nil; curNode = curNode.prev {
		curNode.next, curNode.prev = curNode.prev, curNode.next
	}

	list.head, list.tail = list.tail, list.head
}
