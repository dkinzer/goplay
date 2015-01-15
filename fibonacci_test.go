package main

import (
  "testing"
)

func TestFibonnacci(t *testing.T) {
  f := fibonacci()
  assert(f() == 0, "f0 == 0", t)
  assert(f() == 1, "f1 == 1", t)
  assert(f() == 1, "f2 == 1", t)
  assert(f() == 2, "f3 == 2", t)
  assert(f() == 3, "f4 == 5", t)
  assert(f() == 5, "f6 == 8", t)
}
