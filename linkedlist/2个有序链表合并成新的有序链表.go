package linkedlist

// TwoLinkedListAddition 2个有序链表合并成一个新的有序链表
func TwoLinkedListAddition(list1, list2 *SingleLinkedList) *SingleLinkedList {
	//len1 := list1.Len()
	//len2 := list2.Len()
	ret := new(SingleLinkedList)
	current1, current2 := list1.Node, list2.Node
	for current1 != nil && current2 != nil {
		if current1.Data.(int) < current2.Data.(int) {
			if ret.Node == nil {
				ret.Node = current1

			} else {
				ret.Node.next = current1

			}

			current1 = current1.next
		} else {
			ret.Node = current2
			current2 = current2.next
		}
	}

	if current1 != nil {
		ret.Node = current1
		current1 = current1.next
	}

	if current2 != nil {
		ret.Node = current2
		current2 = current2.next
	}

	return ret
}
