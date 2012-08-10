package main

func main() {

  println(change())

}

func change() (n int) {

  n = 10

  // Proof that defered functions get called AFTER
  // the value passed to 'return' is evaluated
  // Kinda weird fo' sho'...
  defer func() {
    n++
  }()

  return 5
}
