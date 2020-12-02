package lib

import (
	"strconv"
	"log"
)

func ArrayToInt(array []string) []int {
	var intArray []int
	for i, it := range array {
		item, err := strconv.Atoi(it)
		intArray = append(intArray, item)
		if err != nil {
			log.Fatalf("readLines: %s", i, err)
		}
	}
	return intArray
}