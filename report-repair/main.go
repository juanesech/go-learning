package main

import (
  "fmt"
  "log"
  L "./lib"
)

func main()  {
  var sum = 2020
  input_content, err := L.ReadLines("input.txt")
  input := L.ArrayToInt(input_content)
  result_i2 := L.Iterator(input, 2, sum)
  result_i3 := L.Iterator(input, 3, sum)
  if err != nil {
      log.Fatalf("readLines: %s", err)
  }
  fmt.Println("2 Numbers", result_i2)
  fmt.Println("3 Numbers", result_i3)
}
