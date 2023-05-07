package linkedlist

import "fmt"

type SingleLinkedNode struct {
	data interface{}
	next *SingleLinkedNode
}

type SingleLinkedList struct {
	head *SingleLinkedNode  //指向第一个节点的指针
}


func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		head: nil,
	}
}

func (sl *SingleLinkedList) Add(data interface{}) {
	node := &SingleLinkedNode{
		data: data,
	}

	if sl.head == nil {
		sl.head = node
	}else {
		current := sl.head
		for current.next != nil {
			current = current.next
		}

		current.next = node
	}
}

func (sl *SingleLinkedList) List()  {
	if sl.head == nil {
		fmt.Println("链表为空!!!")
		return
	}
	current := sl.head
	for current.next != nil {
		fmt.Println("data:", current.data)
		current = current.next
	}
	fmt.Println("data:", current.data)
}

func (sl *SingleLinkedList) Delete(target interface{})   {
	if sl.head == nil {
		fmt.Println("链表为空!!!")
		return
	}

	//假如是第一个节点
	if sl.head.data == target {
		sl.head = sl.head.next
		return
	}

	current := sl.head
	for current.next != nil {
		//用current.next.data 因为当前的节点无法找到前一个节点，也无法将前一个节点连上后一个节点
		if current.next.data == target {
			//找到要删除的目标
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

//Len 返回链表的长度
func (sl *SingleLinkedList)Len() int {
	if sl.head == nil {
		return 0
	}
	len := 1
	current := sl.head
	for current.next != nil {
		len++
		current = current.next
	}
	return len
}
