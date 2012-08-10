package main

import "fmt"

func main() {

  fmt.Printf("%T\n", 1)
  fmt.Printf("%T\n", 'π')
  fmt.Printf("%T\n", "π")
  fmt.Printf("%T\n", [...]int{1})
  fmt.Printf("%T\n", [[]int]int{{1}, {2}})
  fmt.Printf("%T\n", [...]rune{'ç'})
  fmt.Printf("%T\n", func(){})
}
