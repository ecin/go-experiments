package main

import "fmt"

func main() {

  xs := []int{1, 3, 2, 4, 6, 5, 0, 100, -234, 222}
  fmt.Printf("%v\n", bubbleSort(xs))

}

// Actually don't need to return xs here; slices are passed around
// as references, not copied
func bubbleSort(xs []int) []int {
  for {
    swapped := false
    for i := 0; i < len(xs) - 1; i++ {
      if xs[i] > xs[i+1] {
        xs[i], xs[i+1] = xs[i+1], xs[i]
        swapped = true
      }
    }
    if (!swapped) {
      break
    }
  }

  return xs
}
