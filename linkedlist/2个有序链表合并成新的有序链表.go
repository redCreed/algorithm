package linkedlist

// TwoLinkedListAddition 2个有序链表合并成一个新的有序链表
func TwoLinkedListAddition(list1, list2 *SingleLinkedList) *SingleLinkedList {
	ret := new(SingleLinkedList)
	current1, current2 := list1.Node, list2.Node
	temp := new(SingleLinkedNode)
	ret.Node = temp
	for current1 != nil && current2 != nil {
		if current1.Data.(int) < current2.Data.(int) {
			temp.next = current1
			current1 = current1.next
		} else {
			temp.next = current2
			current2 = current2.next
		}

		temp = temp.next
	}

	if current1 != nil {
		temp.next = current1
	}

	if current2 != nil {
		temp.next = current2
	}

	return ret
}
