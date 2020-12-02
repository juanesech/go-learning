package lib

import (
	"strconv"
	"log"
)


//ArrayToInt function Convert string array to int one
func ArrayToInt(array []string) []int {
	var intArray []int
	for i := range array {
		item, err := strconv.Atoi(array[i])
		intArray = append(intArray, item)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
	}
	return intArray
}