package utils

// Returns true if two string slices contain the same elements (regardless of order)
func EqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[string]int)
	for _, v := range a {
		counts[v]++
	}
	for _, v := range b {
		counts[v]--
	}
	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}
