package main

import(
	"fmt"
)

type Stack struct {
	x []int
}

func (s *Stack) Push(no int) {
	s.x = append(s.x, no)
}

// Peek function panics for empty stack
func (s *Stack) Peek() int {
	return s.x[len(s.x)-1]
}

// Pop function panics for empty stack
func (s *Stack) Pop() int {
	no := s.Peek()
	s.x = s.x[:len(s.x)-1]
	return no
}

func main() {
	s := new(Stack)

	s.Push(1)
	s.Push(2)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
