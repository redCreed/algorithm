package queue

import (
	"fmt"
	"testing"
)

func TestNewLoopQueue(t *testing.T) {
	l := NewLoopQueue(4) //最大有效长度是3
	if err := l.Push(2); err != nil {
		fmt.Println("err:", err)
	}
	if err := l.Push(3); err != nil {
		fmt.Println("err:", err)
	}

	if err := l.Push(4); err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(fmt.Sprintf("%+v", l))
	if value, err := l.Pop(); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("value:", value)
	}

	fmt.Println(fmt.Sprintf("%+v", l))
}
