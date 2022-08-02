package stack

type Stack struct {
	items []int
	top   int
}

func (s *Stack) Create(max int) []int {
	s.items = make([]int, 0, max)
	s.top = -1
	return s.items
}

func (s *Stack) IsFull() bool {
	return s.top == cap(s.items)-1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Push(v int) {
	if s.IsFull() {
		return
	} else {
		s.top += 1
		s.items = append(s.items, v)
		return
	}
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	} else {
		s.items = s.items[:s.top]
		s.top -= 1
	}
}

func (s *Stack) Peek() int {
	return s.items[s.top]
}
