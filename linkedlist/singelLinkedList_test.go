package linkedlist

import (
	"fmt"
	"testing"
)

func TestNewSingleLinkedList(t *testing.T) {
	sl := NewSingleLinkedList()
	sl.Add("张三")
	sl.Add("李四")
	sl.Add("33")
	sl.List()
	sl.Delete("33")
	fmt.Println("delete:")
	sl.List()
	fmt.Println("len:", sl.Len())

	//测试倒数第k个节点
	fmt.Println("node:", fmt.Sprintf("%+v", sl.GetIndexKNode(1)))

}
