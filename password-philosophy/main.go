package main

import (
	"fmt"
	"log"

	U "../utils"
	L "./lib"
)

func main() {
	fileContent, err := U.ReadLines("input.txt")
	var counterV1 int
	var counterV2 int
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, value := range fileContent {
		formatedValue := L.Format(value)
		if L.PassValidator(formatedValue) {
			counterV1++
		}
		if L.PassValidatorV2(formatedValue) {
			counterV2++
		}
	}
	fmt.Println("Count V1: ", counterV1)
	fmt.Println("Count V2: ", counterV2)
}
