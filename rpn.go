package main

import "stack"

func main() {
  s := new(stack)
  s.Push(5)
  s.Push(10)
  println(s.Pop())
}
