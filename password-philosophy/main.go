package main

import (
	"fmt"
	"log"

	U "../utils"
	L "./lib"
)

func main() {
	fileContent, err := U.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, value := range fileContent {
		formatedValue := L.Format(value)
		fmt.Println(formatedValue)
	}
}
