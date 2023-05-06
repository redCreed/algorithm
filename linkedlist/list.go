package linkedlist

type LinkedList struct {
	data interface{}
	next *LinkedList
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		data: nil,
		next: nil,
	}
}

func (l *LinkedList) Add(data interface{}) {

}
