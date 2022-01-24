package ds

import "fmt"

/*Stack*/
type Stack struct{ data []int }

/*creates new stacks*/
func NewStack() Stack { return Stack{data: []int{}} }

/*gives length of stack*/
func (s *Stack) Len() int { return len(s.data) }

/*push value in stack*/
func (s *Stack) Push(val int) { s.data = append(s.data, val) }

/*pop value from stack*/
func (s *Stack) Pop() (ans int) {
	ans = s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return
}

/*peek current value in stack*/
func (s *Stack) Peek() int { return s.data[s.Len()-1] }

/*Checks if stack is empty or not*/
func (s *Stack) IsEmpty() bool { return s.Len() == 0 }

/*Print the stack on standard output*/
func (s *Stack) Show() { fmt.Println(s.data) }
