package stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack(4)
	if err := stack.Push(1); err != nil {
		fmt.Println("stack:", 1, err)
	}
	fmt.Println(fmt.Sprintf("%+v", stack))
	if err := stack.Push(2); err != nil {
		fmt.Println("stack:", 2, err)
	}
	fmt.Println(fmt.Sprintf("%+v", stack))
	if value, err := stack.Pop(); err != nil {
		fmt.Println("stack:", 2, err)
	} else {
		fmt.Println("value:", value)
	}
	fmt.Println(fmt.Sprintf("%+v", stack))
}

