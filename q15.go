// Functions that return functions are fun and easy!

package main

func main() {

  f := plusTwo()
  println(f(0))
  println(f(2))
  println(f(4))

  f = plusX(-1)
  println(f(0))
  println(f(5))
  println(f(10))

}

func plusTwo() func(int)(int) {
  return func(x int) int {
    return x + 2
  }
}

func plusX(plus int) func(int)(int) {
  return func(x int) int {
    return x + plus
  }
}
