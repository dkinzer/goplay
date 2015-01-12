package main

import (
  "fmt"
)

type address struct {
  street1 string
  street2 string
  city string
  state string
  zipcode int
}

var myAddress = address {
  "510 Spruce Street",
  "Apt. 1R", 
  "Philadelphia", 
  "PA",
  19106,
}

func smallIntArray () [3]int {
  var array [3]int
  return array
}

func incFirstElement (a *[3]int) {
  a[0] = a[0] + 1
}


func incFirstElementSlice (a []int) {
  a[0] = a[0] + 1
}

// Slices are pretty dangerous.
// This operation will change the original array.
func head (xs []int) int {
  xs[1] = 100
  return xs[0]
}

func rest (xs []int) []int {
  return xs[1:]
}

func zeroIt (x *int64) {
  *x = *x * 0
}

func transplant (address *address, city string) {
  // Pointer inderection is transparent for structs (i.e. no * needed).
  address.state = city
}

// Making slices.
func makeR (c int) []int {
  return make([]int, 0, c)
}

func forRange (s []int) {
  for _, v := range s {
    fmt.Printf("2 times %d = %d\n", v, (2 * v))
  }
}

func moreTypes() {
  // Strucs
  fmt.Println("Hello Strucs!:")
  fmt.Println(myAddress)
  fmt.Println(myAddress.city)

  // Pointers
  ten := int64(10)
  fmt.Println("Hello Pointers!:")
  zeroIt(&ten)
  fmt.Println(ten)

  // Pointers to Strucs
  transplant(&myAddress, "New York City")
  fmt.Println(myAddress)

  // Arrays.
  sarray := smallIntArray()
  incFirstElement(&sarray)
  fmt.Println(sarray)

  // Slices.
  slarray := []int {1, 2, 3, 4, 5}
  fmt.Println("Hello Slices!:")
  incFirstElementSlice(slarray)
  fmt.Println(slarray)
  fmt.Println(head(slarray))
  fmt.Println(rest(slarray))
  fmt.Println(slarray)

  s := makeR(20)
  fmt.Println(append(s, 1, 2, 3, 4, 5, 6))
  fmt.Println(s, len(s), cap(s))

  //Ranges
  forRange(append(s, 10, 2, 30, 4, 50, 6))
}
