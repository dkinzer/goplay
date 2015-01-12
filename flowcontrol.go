package main

import (
  "math"
  "time"
  "fmt"
)

func forr (x int) int {
  for i := 0; i < 100; i++ {
    x += i
  }
  return x
}

func forall (x int) int {
  sum := x
  for sum < 1000 {
    sum += sum
  }
  return sum
}

func greaterThan (x, y int) bool {
  if x > y {
    return true
  } 
  return false
}

// variables inside of if statments are also availabe indie of blocks.
func ifelse () string {
  if name := "Olvia"; false {
  } else {
    return name
  }
  return "Wrong"
}

func Sqrt(x float64) float64 {
  sqrt := float64(3)
  delta := float64(1)

  for delta > 0.0000000000001 {
    z := sqrt
    sqrt = z - ((z * z) - x) / (2 * z)
    delta = math.Abs(sqrt - z)
  }

  return sqrt
}

// case body ends automatically unless it ends with a fallthrough
func greet (name string) string {
  var reply string

  switch name {
    case "Olivia":
      reply = "Hey! How was your sleep. Did you have any nice dreams?"
    case "Kirsten":
      reply = "Hey K! How was your day?"
    default:
      reply = "Have a good one!"
  }
  return reply
}

// use switch without a condition for long if then elses.
func sleep() string {
  var reply string;
  t := time.Now()
  switch {
    case t.Hour() > 23:
      reply = "You should be asleep."
    case t.Hour() > 9:
      reply = "You should be at work"
    case t.Hour() > 17:
      reply = "You should go home."
  }
  return reply
}

func helloWorld() {
  defer fmt.Println("World.")
  fmt.Println("Hello")
}


// Defers are stacked; last in is first out.
func stackDefer() {
  for i := 0 ; i < 10 ; i ++ {
    defer fmt.Println(i)
  }
  fmt.Println("done")
}

func flowControls() {
  fmt.Println (ifelse())
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(2) * Sqrt(2))
  fmt.Println(greet("Olivia"))
  fmt.Println(greet("Kirsten"))
  fmt.Println(greet("David"))
  fmt.Println(sleep())
  helloWorld()
  stackDefer()
}
