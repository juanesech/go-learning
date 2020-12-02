package lib

import (
	"fmt"
)

func Iterator(in []int, iter int, sumVal int) []int {
	var values []int
	switch iter {
	case 2:
		for i := range in {
			for a := range in {
			  if in[i]+in[a] == sumVal {
				values = append(values, in[i], in[a])
				fmt.Println("Result: ", in[i]*in[a])
				return values
			  }
			}
		}
	case 3:
		for i := range in {
			for a := range in {
				for b := range in {
					if in[i]+in[a]+in[b] == sumVal {
						values = append(values, in[i], in[a], in[b])
						fmt.Println("Result: ", in[i]*in[a]*in[b])
						return values
					  }
				}
			}
		}
	}
	return values
}
