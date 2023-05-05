package stack

import "errors"

const (
	defaultMaxSize = 1000
)

type Stack struct {
	maxSize int
	top     int
	data    []int
}

func NewStack(size int) *Stack {
	if size <= 0 {
		size = defaultMaxSize
	}

	return &Stack{
		maxSize: size,
		top:     -1,
		data:    make([]int, size, size),
	}
}

// isFull 判断队列是否已经满
func (s *Stack) isFull() bool {

	return s.top == s.maxSize-1
}

// isFull 判断队列是否为空
func (s *Stack) isEmpty() bool {

	return s.top == -1
}

// Push 向队列添加一个数据
func (s *Stack) Push(d int) error {
	//判断队列是否已经满
	if s.isFull() {
		return errors.New("the stack is already full and cannot be pushed")
	}
	s.top++
	s.data[s.top] = d

	return nil
}

// Pop 向队列获取一个数据
func (s *Stack) Pop() (int, error) {
	//判断队列是否为空
	if s.isEmpty() {
		return 0, errors.New("the stack is empty")
	}
	value := s.data[s.top]
	s.data[s.top] = 0
	s.top--
	return value, nil
}
