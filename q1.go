package main

import "fmt"

func main() {

  /*
  for i := 0; i < 10; i++ {
    fmt.Println(i);
  }
  */

  /*
  i := 0

  I: fmt.Println(i)

  i++
  if i < 10 {
    goto I
  }
  */

  a := [10]int{}

  for i := 0; i < 10; i++ {
    a[i] = i;
  }

  fmt.Printf("%v", a);
}


