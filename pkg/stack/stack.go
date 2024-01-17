package stack

type Stack[T any] struct {
	Len  int
	head *sNode[T]
}

type sNode[T any] struct {
	data T
	next *sNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{Len: 0, head: nil}
}

func (s *Stack[T]) Peek() (data T) {
	if s.Len <= 0 {
		return
	}
	return s.head.data
}

func (s *Stack[T]) Pop() (data T) {
	if s.Len <= 0 {
		return
	}
	s.Len--
	if s.head == nil {
		return
	}
	data = s.head.data
	s.head = s.head.next
	return
}

func (s *Stack[T]) Push(data T) {
	s.Len++
	newNode := &sNode[T]{data, nil}
	if s.head == nil {
		s.head = newNode
		return
	}
	newNode.next = s.head
	s.head = newNode
}
