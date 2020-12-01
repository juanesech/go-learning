package main

import (
  "fmt"
  "log"
  L "./lib"
  "strconv"
  "os"
)

func main()  {
  var sum int = 2020
  lines, err := L.ReadLines("input.txt")
  if err != nil {
      log.Fatalf("readLines: %s", err)
  }
  for i, l := range lines {
    line, err := strconv.Atoi(l)
    if err != nil {
        log.Fatalf("readLines: %s", i, err)
    }
    for a, lin := range lines {
      li, err := strconv.Atoi(lin)
      if err != nil {
          log.Fatalf("readLines: %s", a, err)
      }
      if line+li == sum {
        fmt.Println("Values: ", line, li)
        fmt.Println("Result: ", line*li)
        os.Exit(2)
      }
    }
  }
}
