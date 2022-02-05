package filter

type flt func(int) bool

// Filter returns a new slice containing only the elements of the input
func Filter(sl []int, f flt) []int {
	var res []int
	for _, val := range sl {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

// All returns true if all elements of the input satisfy the predicate
func All(sl []int, f flt) bool {
	for _, val := range sl {
		if !f(val) {
			return false
		}
	}
	return true
}

// Any returns true if any element of the input satisfies the predicate
func Any(sl []int, f flt) bool {
	for _, val := range sl {
		if f(val) {
			return true
		}
	}
	return false
}
