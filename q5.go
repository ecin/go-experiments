package main

import "fmt"

func main() {

  // Doesn't seem like the ellipsis is necessary
  floats := [...]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  slice := floats[:]

  var avg float64

  for _,n := range slice {
    avg += n / 10.0
  }

  fmt.Println(avg)


}:e
