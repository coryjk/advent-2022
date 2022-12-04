package functional

func Max[V comparable](arr []V, predicate func(left V, right V) bool) V {
	var max V
	for i, value := range arr {
		if i == 0 {
			max = value
		} else {
			// re-assign new max value if predicate passes
			if predicate(value, max) {
				max = value
			}
		}
	}
	return max
}
