package linkedlist

/*
    查找单链表的反转
	https://www.cnblogs.com/wll500/p/16202749.html
	https://blog.csdn.net/weixin_35753431/article/details/128870048
*/
func (sl *SingleLinkedList) reverse() *SingleLinkedList {
	//空链表
	if sl.head == nil || sl.head.next == nil {
		return sl
	}

	current := sl.head
	var preNode *SingleLinkedNode
	temp := new(SingleLinkedNode)
	for current != nil {
		//保存current的下一个节点
		temp = current.next
		current.next = preNode
		preNode = current
		current = temp
	}

	return &SingleLinkedList{head: preNode}
}

//Node is a single element in a linked list
type Node struct {
	Val  int
	Next *Node
}

// Reverse reverses a linked list
func Reverse(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
