package strings

import "sort"

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortRunes) Less(i int, j int) bool {
	return s[i] < s[j]
}

// SortString: Sort the given string.
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
