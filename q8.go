package main

import "fmt"

type stack struct {
  i int
  data [10]int
}

func main() {

  var s = new(stack)
  s.push(1).push(2).push(3)

  fmt.Printf("%v\n", s)
  fmt.Printf("%d %d %d\n", s.pop(), s.pop(), s.pop())

}

func (s *stack) push(n int) *stack {

  if s.i < len(s.data) {
    s.data[s.i] = n
    s.i += 1
  }

  return s
}

// Return value needs parens when named
func (s *stack) pop() (value int) {

  if s.i == 0 {
    return nil
  }

  if s.i > 0 {
    s.i -= 1
  }

  value = s.data[s.i]

  return

}

// Interface for printing with %v
func (s *stack) Stringer() string {
  return "IMPLEMENT ME!"
}
