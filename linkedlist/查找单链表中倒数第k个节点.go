package linkedlist



//查找单链表中倒数第k个节点数, 默认节点index从0开始
func (sl *SingleLinkedList) GetIndexKNode(k int) *SingleLinkedNode {
	//获取链表长度，那倒数第k个节点的index为len-k
	len := sl.Len()
	if len< k {
		return nil
	}
	kNodeIndex := len-k
	if kNodeIndex == 0 {
		return sl.head
	}
	current := sl.head
	index := 0
	for current.next != nil {
		index ++
		if index == kNodeIndex {
			return current.next
		}
		current = current.next
	}

	return nil
}
