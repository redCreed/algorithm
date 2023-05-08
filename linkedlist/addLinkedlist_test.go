package linkedlist

import (
	"fmt"
	"testing"
)

func TestTwoLinkedListAddition(t *testing.T) {
	l1 := new(SingleLinkedList)
	l2 := new(SingleLinkedList)

	l1.Add(1)
	l1.Add(3)
	l1.Add(5)
	l1.Add(6)

	l2.Add(2)
	l2.Add(4)
	fmt.Println("list1:")
	l1.List()
	fmt.Println("list2:")
	l2.List()

	fmt.Println("add list1 and list2:")
	TwoLinkedListAddition(l1, l2).List()
}
