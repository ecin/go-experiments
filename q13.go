package main

import "math"

func main() {

  xs := []int{7, 2, 3, 4, 5, 6}
  // int(Inf+) actually results in the smallest int possible. =/
  println(int(math.Inf(0)))
  println(max(xs))
  println(min(xs))

}

func pairMax(x int, y int) (m int) {
  m = x

  if y > x {
    m = y
  }

  return
}

func pairMin(x int, y int) (m int) {
  m = x
  if x > y {
    m = y
  }

  return
}

func max(xs []int) (m int) {
  m = int(math.Inf(-1))
  for _, x := range xs {
    m = pairMax(m, x)
  }

  return
}

func min(xs []int) (m int) {
  m = xs[0]
  for _, x := range xs {
    m = pairMin(m, x)
  }

  return
}
