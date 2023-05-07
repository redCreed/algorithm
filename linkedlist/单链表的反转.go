package linkedlist

/*
    查找单链表的反转
	https://www.cnblogs.com/wll500/p/16202749.html
*/
func (sl *SingleLinkedList) reverse() *SingleLinkedList{
	//空链表
	if sl.head == nil || sl.head.next == nil {
		return sl
	}

	current := sl.head
	newSl :=new(SingleLinkedList)
	newHead := new(SingleLinkedNode)
	temp := new(SingleLinkedNode)
	for current.next != nil {
		//保存current的下一个节点
		temp = current.next
		current.next = newHead.next

		current = temp
	}
	return newSl
}

