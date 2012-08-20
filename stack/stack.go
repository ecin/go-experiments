package Stack

type Stack struct {
  Size int
  elements [10]int
}

func (s *Stack) Push(x int) Stack {
  s.elements[s.Size] = x
  s.Size++

  return *s
}

func (s *Stack) Pop() (x int) {
  if (s.Size > 0) {
    s.Size--
    x = s.elements[s.Size]
  }

  return
}

func (s *Stack) Peek() int {
  return s.elements[s.Size - 1]
}
