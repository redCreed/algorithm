package linkedlist

/*
	    查找单链表的反转
		https://www.cnblogs.com/wll500/p/16202749.html
		https://blog.csdn.net/weixin_35753431/article/details/128870048
*/
func (sl *SingleLinkedList) reverse() *SingleLinkedList {
	//空链表
	if sl.Node == nil || sl.Node.next == nil {
		return sl
	}

	current := sl.Node
	var preNode *SingleLinkedNode
	temp := new(SingleLinkedNode)
	for current != nil {
		//保存current的下一个节点
		temp = current.next
		current.next = preNode
		preNode = current
		current = temp
	}

	return &SingleLinkedList{Node: preNode}
}

// Node is a single element in a linked list
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

func Reverse1(node *Node) *Node {
	var ret *Node
	current := node
	for current != nil {
		//下一个对象
		next := current.Next
		current.Next = ret
		ret = current
		//把旧链表的下一个节点复制给current
		current = next
	}
	return node
}

func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		// 保证 0 ~ i-1 有序
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		// 填坑 插入位置
		nums[j+1] = key
	}
	return nums
}
