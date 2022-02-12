package slice

//Max return max value of a slice
func Max(args ...int) int {
	lm := args[0]
	for _, arg := range args {
		if arg > lm {
			lm = arg
		}
	}
	return lm
}

//Min return min value of a slice
func Min(args ...int) int {
	lm := args[0]
	for _, arg := range args {
		if arg < lm {
			lm = arg
		}
	}
	return lm
}
