package main

func main() {

  printInts(1, 2, 3, 4, 5)

}

func printInts(xs ...int) {
  for _, n := range xs {
    println(n)
  }
}
