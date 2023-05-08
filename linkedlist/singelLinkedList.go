package linkedlist

import "fmt"

type SingleLinkedNode struct {
	Data interface{}
	next *SingleLinkedNode
}

type SingleLinkedList struct {
	Node *SingleLinkedNode //指向第一个节点的指针
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		Node: nil,
	}
}

func (sl *SingleLinkedList) Add(data interface{}) {
	node := &SingleLinkedNode{
		Data: data,
	}

	if sl.Node == nil {
		sl.Node = node
	} else {
		current := sl.Node
		for current.next != nil {
			current = current.next
		}

		current.next = node
	}
}

func (sl *SingleLinkedList) List() {
	if sl.Node == nil {
		fmt.Println("链表为空!!!")
		return
	}
	current := sl.Node
	for current.next != nil {
		fmt.Println("data:", current.Data)
		current = current.next
	}
	fmt.Println("data:", current.Data)
}

func (sl *SingleLinkedList) Delete(target interface{}) {
	if sl.Node == nil {
		fmt.Println("链表为空!!!")
		return
	}

	//假如是第一个节点
	if sl.Node.Data == target {
		sl.Node = sl.Node.next
		return
	}

	current := sl.Node
	for current.next != nil {
		//用current.next.data 因为当前的节点无法找到前一个节点，也无法将前一个节点连上后一个节点
		if current.next.Data == target {
			//找到要删除的目标
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

//Len 返回链表的长度
func (sl *SingleLinkedList) Len() int {
	if sl.Node == nil {
		return 0
	}
	len := 1
	current := sl.Node
	for current.next != nil {
		len++
		current = current.next
	}
	return len
}
