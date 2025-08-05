package utils

import (
	"reflect"
	"sort"
	"strings"
)

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

func MultipleCheck(a, b []string) bool {
	normalize := func(arr []string) []string {
		seen := make(map[string]struct{})
		var result []string
		for _, v := range arr {
			normalized := strings.TrimSpace(strings.ToLower(v))
			if _, exists := seen[normalized]; !exists {
				seen[normalized] = struct{}{}
				result = append(result, normalized)
			}
		}
		sort.Strings(result)
		return result
	}

	aNorm := normalize(a)
	bNorm := normalize(b)

	return reflect.DeepEqual(aNorm, bNorm)
}
