package algorithms

import "errors"

// Stack, along with methods Push and Pop implements a LIFO data structure.
type stack struct {
	items []int
	index int /* the next element in items */
}

// push adds an int to the stack.
func (s *stack) push(i int) {
	if len(s.items) == s.index { /* s.items[s.index] is out of bounds. */
		s.grow()
	}
	s.items[s.index] = i
	s.index++
}

// pop returns the last item from a stack.items, and removes it.
func (s *stack) pop() (int, error) {
	if s.index == 0 {
		return 0, errors.New("pop: cannot pop empty stack")
	}
	if len(s.items) == s.index || s.index > 0 { /* s.items[s.index] is out of bounds. */
		s.index--
	}
	item := s.items[s.index]
	return item, nil
}

// grow extends the size of s.items[].
func (s *stack) grow() {
	var newA []int
	if len(s.items) == 0 {
		newA = make([]int, 1)
	} else {
		newA = make([]int, len(s.items)*2)
	}
	for i, v := range s.items {
		newA[i] = v
	}
	s.items = newA
}
