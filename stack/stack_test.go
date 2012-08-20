package Stack

import "testing"

func TestPush(t *testing.T) {
  s := new(Stack)

  if (s.Size != 0) {
    t.Error("Size of a new Stack should be 0")
  }

  s.Push(1)

  if (s.Size != 1) {
    t.Error("Size of a Stack should increase by 1 after a push")
  }
}

func TestPop(t *testing.T) {
  s := new(Stack)

  s.Push(9)
  if (s.Pop() != 9) {
    t.Error("Stack.Pop() should return last pushed element")
  }

  if (s.Size != 0) {
    t.Error("Size of a Stack should decrease by 1 after a pop")
  }
}

func TestPeek(t *testing.T) {
  s := new(Stack)

  s.Push(9)
  if (s.Peek() != 9) {
    t.Error("Stack.Peek() should return last pushed element")
  }

  if (s.Peek() != 9) {
    t.Error("Stack.Peek() should always return the same value in consecutive calls")
  }

  if (s.Size != 1) {
    t.Error("Stack.Peek() should not decrease size of Stack")
  }
}
