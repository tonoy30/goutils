package math

// Max returns the largest of x, y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the min value of x, y.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
