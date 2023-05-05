package queue

import (
	"errors"
	"fmt"
)

/*
	数组做循环队列
	参考文档:
	https://cloud.tencent.com/developer/article/2020905
	https://blog.csdn.net/weixin_42117918/article/details/81839751
*/

const (
	defaultLoopMaxSize = 1000
)

type LoopQueue struct {
	front   int   //头元素下标
	rear    int   //尾元素下标
	maxSize int   //队列最大长度
	data    []int //模拟用int类型切片
}

func NewLoopQueue(size int) *LoopQueue {
	if size <= 0 {
		size = defaultLoopMaxSize
	}

	return &LoopQueue{
		maxSize: size,
		data:    make([]int, size, size),
	}
}

// isFull 判断队列是否已经满
func (lq *LoopQueue) isFull() bool {
	//todo 这样判断会导致队列少一个位置
	return (lq.rear+1)%lq.maxSize == lq.front
}

// isFull 判断队列是否为空
func (lq *LoopQueue) isEmpty() bool {
	return lq.rear == lq.front
}

// Push 向队列添加一个数据
func (lq *LoopQueue) Push(d int) error {
	//判断队列是否已经满
	if lq.isFull() {
		return errors.New("the queue is already full and cannot be pushed")
	}
	lq.data[lq.rear] = d
	lq.rear = (lq.rear + 1) % lq.maxSize
	return nil
}

// Pop 向队列获取一个数据
func (lq *LoopQueue) Pop() (int, error) {
	//判断队列是否为空
	if lq.isEmpty() {
		return 0, errors.New("the queue is empty")
	}
	value := lq.data[lq.front]
	//清空该位置数据
	lq.data[lq.front] = 0
	lq.front = (lq.front + 1) % lq.maxSize

	return value, nil
}

// Len 获取队列长度
func (lq *LoopQueue) Len() int {

	return (lq.rear - lq.front + lq.maxSize) % lq.maxSize
}

func (lq *LoopQueue) Data() []int {
	return lq.data
}

// ShowQueue 打印信息
func (lq *LoopQueue) ShowQueue() {
	for i := lq.front; i < lq.front+lq.Len(); i++ {
		fmt.Println("data:", i%lq.maxSize, lq.data[i%lq.maxSize])
	}
}
