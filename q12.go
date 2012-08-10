// Passing functions around as parameters

package main

import "fmt"

func main() {

  xs := []int{1, 2, 3, 4, 5,}

  fmt.Println("%v", xs)
  fmt.Println("%v", mapFunc(xs, addOne))
  fmt.Println("%v", mapFunc(xs, double))

}

func addOne(x int) int {
  return x + 1
}

func double(x int) int {
  return 2 * x
}

func mapFunc(xs []int, f func(int)(int)) []int {
  for pos, x := range xs {
    xs[pos] = f(x)
  }

  return xs
}
