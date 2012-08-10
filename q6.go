package main

import "fmt"

func main() {

  floats := []float64{1, 2, 3, 4, 5.631, 7.8}

  fmt.Printf("%f\n", average(floats))

}

func average(floats []float64) float64 {
  var sum float64

  for _, n := range floats {
    fmt.Println(n)
    sum += float64(n)
  }

  return sum / float64(len(floats))
}
