package main

func main() {

  for i := 0; i < 10; i++ {
    println(fib(i))
  }

}

func fib(i int) (value int) {
  switch i {
  case 0:
    value = 0
  case 1:
    value = 1
  default:
    value = fib(i-1) + fib(i-2)
  }

  return
}
