package linkedlist

// TwoLinkedListAddition 2个有序链表合并成一个新的有序链表
func TwoLinkedListAddition(list1, list2 *SingleLinkedList) *SingleLinkedList {
	//len1 := list1.Len()
	//len2 := list2.Len()
	ret := new(SingleLinkedList)
	current1, current2 := list1.Node, list2.Node
	var temp *SingleLinkedNode
	for current1 != nil && current2 != nil {
		if current1.Data.(int) < current2.Data.(int) {
			//if temp == nil {
			//	temp = current1
			//
			//} else {
			temp.next = current1
			//}

			current1 = current1.next
		} else {
			//if temp == nil {
			//	temp = current2
			//
			//} else {
			temp.next = current2
			//}

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

	ret.Node = temp
	return ret
}

// 遍历 l1 和 l2，将较小的节点逐个加入到合并后的链表中
//for l1 != nil && l2 != nil {
//if l1.Val > l2.Val {
//cur.Next = l2
//l2 = l2.Next
//} else {
//cur.Next = l1
//l1 = l1.Next
//}
//// 移动 cur 指针到下一个节点
//cur = cur.Next
//}
//
//// 如果 l1 还有剩余节点，将剩余节点全部加入到合并后的链表中
//if l1 != nil {
//cur.Next = l1
//} else if l2 != nil { // 如果 l2 还有剩余节点，将剩余节点全部加入到合并后的链表中
//cur.Next = l2
//}
