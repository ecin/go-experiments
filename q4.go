package main

import "fmt"

func main() {

  /*
  for i := 1; i <= 100; i++ {
    for j := 0; j < i; j++ {
      fmt.Print("A");
    }
    fmt.Println();
  }
  */

  a := "asSASA ddd dsjkdsjs çç"
  chars, bytes := 0, 0

  // Range iterates over the runes, e.g. characters
  // in the string
  // Also tells us the position of each rune, which we
  // can use to see how many bytes the string has
  for pos,_ := range a {
    chars++
    bytes += pos
  }

  fmt.Print("Number of characters in " + a + " is: ")
  fmt.Println(chars)
  fmt.Print("Number of bytes in " + a + " is: ")
  fmt.Println(bytes)

  fmt.Println(reverse("hello world"))

  runes := []rune(a)

  for pos,_ := range runes {
    switch pos {
      case 4: runes[pos] = 'a'
      case 5: runes[pos] = 'b'
      case 6: runes[pos] = 'c'
    }
  }

  fmt.Println(string(runes))

}

func reverse(a string) string {
  r := []rune(a)

  for i, j := len(a) - 1, 0; i >= len(a) / 2; i, j = i - 1, j + 1 {
    r[i], r[j] = r[j], r[i]
  }

  return string(r);
}
