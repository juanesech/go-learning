package main

import (
	"fmt"
	"log"

	U "../utils"
	L "./lib"
)

func main() {
	fileContent, err := U.ReadLines("input.txt")
	var counter int
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, value := range fileContent {
		formatedValue := L.Format(value)
		if L.PassValidator(formatedValue) {
			counter++
		}
	}
	fmt.Println("Count: ", counter)
}
