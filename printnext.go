package main

func main() {

  // Type coercion example. '9' gets translated to its
  // ASCII value; works for UTF-8 chars as well
  //println(60 - '9')
  var q, w int
  println(q)
  println(w)
  println("-")

  a := []byte{'1', '2', '3', '4'}
  var x int
  for i := 0; i < len(a); {
    x, i = nextInt(a, i)
    println(".")
    println(x)
    println(".")
  }

}

func nextInt(b []byte, i int) (int, int) {
  x := 0
  for ; i < len(b); i++ {
    println(x)
    println(int(b[i]))
    println(x*10 + int(b[i]))
    println(x*10 + int(b[i]) - '0')
    x = x*10 + int(b[i]) - '0'
  }
  return x,i
}
