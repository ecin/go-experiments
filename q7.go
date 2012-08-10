// Write a function that returns its (two) parameters in the right, numerical
// (ascending) order

package main

func main() {

  min, max := orderedPair(1, 2)

  println(min)
  println(max)

  min, max = orderedPair(200, 100)

  println(min)
  println(max)

}

func orderedPair(x int, y int) (int, int) {
  if x > y {
    x, y = y, x
  }

  return x, y
}
