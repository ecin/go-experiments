package main

import (
  "bufio"
  "os"
  "fmt"
  "strconv"
)

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

func main() {

  reader := bufio.NewReader(os.Stdin)
  stack := new(Stack)
  var token string

  // Infinite loop
  for {

    input, err := reader.ReadString('\n')
    if err != nil { return }

    // a => index
    // b => UTF-8 value
    for _, char := range input {
      switch {
        case char >= '0' && char <= '9':
          // need to typecast rune to string before concatenating
          token += string(char)
        case char == '\n':
          if (token != "") {
            n, _ := strconv.Atoi(token)
              stack.Push(n)
              token = ""
          }
        case char == '*':
          result := stack.Pop() * stack.Pop()
          fmt.Printf("%d\n", result)
          stack.Push(result)
        case char == '+':
          result := stack.Pop() + stack.Pop()
          fmt.Printf("%d\n", result)
          stack.Push(result)
        case char == '-':
          result := stack.Pop() - stack.Pop()
          fmt.Printf("%d\n", result)
          stack.Push(result)
        case char == 'q':
          return
        default:
      }
    }
  }
}
