// package name can be anything, "main" for standalone executables
package main

import "fmt"

func main() {
// Have to comment this out, 'cause unused vars are a compile error
// (whoa, built-in PMD, boys!)
//  var hello string
//  var times int

//  hello = "Ohayo"
//  times = 5

//  another_number := 1
// The := is useful only to avoid having to type "var"
//  flag := false

  var (
    all, variable, declarations, _ = "all", "variable", "declarations", "discarded"
  )

  const (
    a = iota
    b
    c
  )

  // Does not add a newline at end of string
  fmt.Printf(all + variable + declarations + "\n")
  fmt.Printf("%i\n", a + b + c)
  fmt.Printf(`Testing
              Are spaces considered?`)

  panic
}
