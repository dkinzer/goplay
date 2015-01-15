package main

func fibonacci () func() int {
  fib_m := 0
  fib_n := 1
  return func() int {
    fib := fib_m
    fib_m = fib_n
    fib_n = fib + fib_n
    return fib
  }
}
