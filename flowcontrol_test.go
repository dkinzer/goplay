package main

import "testing"

func assert(truth bool, message string, t *testing.T) {
  if !truth {
    t.Error(message)
  }
}

func TestGreet (t *testing.T) {
  expected := "Hey! How was your sleep. Did you have any nice dreams?" 
  assert(greet("Olivia") == expected, expected, t)
}

func TestForrAll (t *testing.T) {
  assert(forall(1) == 1024, "1024", t)
}

func TestIfElse (t *testing.T) {
  assert(ifelse() == "Olivia", "Olivia", t)
}

func TestGreaterThan(t *testing.T) {
  assert(greaterThan(1,1) == false, "false", t)
  assert(greaterThan(1,2) == false, "false", t)
  assert(greaterThan(2,1) == true, "true", t)
}
