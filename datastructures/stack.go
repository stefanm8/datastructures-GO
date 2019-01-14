package datastructures

type LIFOStack struct {
	stack *LinkedList
}

func (s *LIFOStack) Append(e interface{}) {
	s.stack.Append(e)
}

func (s *LIFOStack) Pop() interface{} {
	e := s.stack.Pop(s.stack.head.data)
	return e
}

func (s *LIFOStack) Len() int {
	return s.stack.Len()
}

func NewStack() *LIFOStack {
	return &LIFOStack{NewLinkedList()}
}
