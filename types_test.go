package main

import "testing"

func TestWordCount(t *testing.T) {
  actual := len(wordCount(""))
  assert(actual == 0, "Empty string returns empty results.", t)

  count := wordCount("hello hello hello world world  !")
  assert(count["hello"] == 3, "3", t)
  assert(count["world"] == 2, "2", t)
  assert(count["something"] == 0, "something doesn't exist", t)
  assert(count["!"] == 1, "!", t)
}
