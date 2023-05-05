package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	front   int   //头元素下标
	rear    int   //尾元素下标
	maxSize int   //队列最大长度
	data    []int //模拟用int类型切片
}

const (
	defaultMaxSize = 1000
)

func NewQueue(size int) *Queue {
	if size <= 0 {
		size = defaultMaxSize
	}

	return &Queue{
		front:   0,
		rear:    -1,
		maxSize: size,
		data:    make([]int, size, size),
	}
}

// isFull 判断队列是否已经满
func (q *Queue) isFull() bool {
	return q.rear+1 == q.maxSize
}

// isFull 判断队列是否为空
func (q *Queue) isEmpty() bool {
	if q.rear-q.front == -1 {
		return true
	}
	return false
}

// Push 向队列添加一个数据
func (q *Queue) Push(d int) error {
	//判断队列是否已经满
	if q.isFull() {
		return errors.New("the queue is already full and cannot be pushed")
	}
	q.rear++
	q.data[q.rear] = d

	return nil
}

// Pop 向队列获取一个数据
func (q *Queue) Pop() (int, error) {
	//判断队列是否为空
	if q.isEmpty() {
		return 0, errors.New("the queue is empty")
	}
	value := q.data[q.front]
	q.data[q.front] = 0
	q.front++
	return value, nil
}

// Len 获取队列长度
func (q *Queue) Len() int {

	return q.rear - q.front + 1
}

func (q *Queue) Data() []int {
	return q.data
}

// ShowQueue 打印信息
func (q *Queue) ShowQueue() {
	for i := q.front; i <= q.Len(); i++ {
		fmt.Println("data:", i, q.data[i])
	}
}
