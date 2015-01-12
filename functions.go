package main

import (
  "fmt"
)

func add (x int, y int) int {
  return x + y
}

func mult (x, y int) int {
  return x * y
}

// functions can return multiple results.
func swap (a, b string) (string, string) {
  return b, a
}

// Named return values act like variables.
func split(sum int) (x, y int) {
  x = sum * 4/9
  y = sum - x
  return
}
// Constant declaration.
const name = "David Kinzer"

// the var statement declares a list of variables.
var c, python, java bool

// some go's types
// Use T(v) to cast v to T

var (
  ints int = 1000000000000000000
  int8s int8 = int8(ints)
  int16s int16 = int16(int8s)
  int32s int32 = int32(int16s)
  int64s int64 = int64(int32s)

  // Type inference.
  aBool = true
)

func functions() {
  const name = "Olivia Kinzer"
  const Big = 1 << 100

  var i int

  // var with initializer
  var c, python, java = true, false, "no!"

  // short variable declarations inside of functions.
  b := java

  fmt.Println( add(100, 30) )
  fmt.Println( mult(100, 30) )
  fmt.Println( swap ("world!", "Hello") )
  fmt.Println( split(100) )
  fmt.Println( i, c, python, java, b, ints, int8s, int16s)
  fmt.Println( name )
  fmt.Println(forr(1))
  fmt.Println(forall(1))
  fmt.Println(greaterThan(3, 4))
}
