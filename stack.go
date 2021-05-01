package stack

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(v interface{}) *Node {
	node := new(Node)
	node.Data = v
	node.Next = nil
	return node
}

type Stack struct {
	stack *Node
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.stack = nil
	return stack
}

func (s *Stack) Push(v interface{}) {
	node := NewNode(v)
	node.Next = s.stack
	s.stack = node
}

func (s *Stack) Pop() interface{} {
	if s.stack != nil {
		aux := s.stack.Data
		s.stack = s.stack.Next
		return aux
	}
	return s.stack.Data
}

func (s *Stack) Peek() interface{} {
	if s.stack == nil {
		return nil
	} else {
		return s.stack.Data
	}
}
