package main

import (
  "fmt"
  "log"
  L "./lib"
)

func main()  {
  var sum = 2020
  fileContent, err := L.ReadLines("input.txt")
  input := L.ArrayToInt(fileContent)
  i2Numbers, i2Product := L.Iterator(input, 2, sum)
  i3Numbers, i3Product := L.Iterator(input, 3, sum)
  if err != nil {
      log.Fatalf("readLines: %s", err)
  }
  fmt.Println("2 Numbers", i2Numbers, "Product: ", i2Product)
  fmt.Println("3 Numbers", i3Numbers, "Product: ", i3Product)
}
