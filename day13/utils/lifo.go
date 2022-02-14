package utils

type Stack struct {
	nodes []string
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n string) {
	if s.count >= len(s.nodes) {
		cnt := len(s.nodes) * 2
		if cnt == 0 {
			cnt = 10
		}
		nodes := make([]string, cnt)
		copy(nodes, s.nodes)
		s.nodes = nodes
	}
	s.nodes[s.count] = n
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() string {
	if s.count == 0 {
		return ""
	}
	node := s.nodes[s.count-1]
	s.count--
	return node
}

func (s *Stack) Len() int {
	return s.count
}
