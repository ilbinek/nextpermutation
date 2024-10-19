package nextpermutation

import (
	"golang.org/x/exp/constraints"
)

func Next[T constraints.Ordered](permutation *[]T, start, end int) bool {
	if end >= len(*permutation) {
		end = len(*permutation) - 1
	}
	
	if start >= end {
		return false
	}

	i := end

	for {
		j := i
		i--

		if (*permutation)[i] < (*permutation)[j] {
			k := end
			for (*permutation)[i] >= (*permutation)[k] {
				k--
			}
			
			iterSwap(permutation, i, k)
			reverse(permutation, j, end)
			return true
		}

		if (i == start) {
			reverse(permutation, start, end);
			return false;
		}
	}
}

func NextString(str *string, start, end int) bool {
	permutation := []rune(*str)
	ret := Next(&permutation, start, end)

	*str = string(permutation)
	return ret
}

func reverse[T comparable](permutation *[]T, start, end int) {
	for start < end {
		(*permutation)[start], (*permutation)[end] = (*permutation)[end], (*permutation)[start]
		start++
		end--
	}
}

func iterSwap[T comparable](permutation *[]T, i, j int) {
	(*permutation)[i], (*permutation)[j] = (*permutation)[j], (*permutation)[i]
}
