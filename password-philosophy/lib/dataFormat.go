package lib

import (
	"log"
	"strconv"
	"strings"
)

//Record struct type
type Record struct {
	min          int
	max          int
	char         string
	recordString string
}

func parseToInt(s string) int {
	intVal, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return intVal
}

//Format function for parse input string to Record struct
func Format(s string) Record {
	splitedString := strings.Split(s, " ")
	stringRange := strings.Split(splitedString[0], "-")
	charString := splitedString[1]
	formatedChar := charString[:len(charString)-1]
	record := Record{
		min:          parseToInt(stringRange[0]),
		max:          parseToInt(stringRange[1]),
		char:         formatedChar,
		recordString: splitedString[len(splitedString)-1]}
	return record
}
