package main

import "fmt"

type curry func() int

func (c curry) chicken() {
  if c() < 1 {
    fmt.Println("Less than one.")
  } else {
    fmt.Println(c())
  }
}


func lemon(i int) curry {
  c := func() int {
    i++
    return 5 + i
  }
  return c
}

