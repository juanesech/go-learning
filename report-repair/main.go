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
  i2Numbers, i2_product := L.Iterator(input, 2, sum)
  i3Numbers, i3_product := L.Iterator(input, 3, sum)
  if err != nil {
      log.Fatalf("readLines: %s", err)
  }
  fmt.Println("2 Numbers", i2Numbers, "Product: ", i2_product)
  fmt.Println("3 Numbers", i3Numbers, "Product: ", i3_product)
}
