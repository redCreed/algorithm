package linkedlist

import (
	"fmt"
	"testing"
)

func TestNewSingleLinkedList1(t *testing.T) {
	sl := NewSingleLinkedList()
	sl.Add("张三")
	sl.Add("李四")
	sl.Add("33")
	sl.List()
 	ret := sl.reverse()
	fmt.Println("node:", fmt.Sprintf("%+v", ret))

}

