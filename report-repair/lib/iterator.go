package lib

func Iterator(in []int, iter int, sumVal int) ([]int, int) {
	var values []int
	var product int
	switch iter {
	case 2:
		for i := range in {
			for a := range in {
			  if in[i]+in[a] == sumVal {
				values = append(values, in[i], in[a])
				product = in[i]*in[a]
				return values, product
			  }
			}
		}
	case 3:
		for i := range in {
			for a := range in {
				for b := range in {
					if in[i]+in[a]+in[b] == sumVal {
						values = append(values, in[i], in[a], in[b])
						product = in[i]*in[a]*in[b]
						return values, product
					  }
				}
			}
		}
	}
	return values, product
}
