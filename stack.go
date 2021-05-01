package stack

//structure Node
type Node struct {
	Data interface{}
	Next *Node
}
//Function to create new node 
func NewNode(v interface{}) *Node {
	node := new(Node)
	node.Data = v
	node.Next = nil
	return node
}

//structure Stack 
type Stack struct {
	stack *Node
}

//Function to create a new stack
func NewStack() *Stack {
	stack := new(Stack)
	stack.stack = nil
	return stack
}

//Method to push node in the stack
func (s *Stack) Push(v interface{}) {
	node := NewNode(v)
	node.Next = s.stack
	s.stack = node
}
//Method to pop node in the stack
func (s *Stack) Pop() interface{} {
	if s.stack != nil {
		aux := s.stack.Data
		s.stack = s.stack.Next
		return aux
	}
	return s.stack.Data
}
//Method to show node top in the stack
func (s *Stack) Peek() interface{} {
	if s.stack == nil {
		return nil
	} else {
		return s.stack.Data
	}
}
