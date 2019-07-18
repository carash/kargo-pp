package kargosort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tc := [][]int{
		[]int{1, 2, 6, 2, 4, 6, 8, 3, 1, 2, 5, 7, 9, 3, 1, 2},
		[]int{6, 4, 3, 4, 7, 9, 0, 4, 2, 3, 4, 32, 3, 5, 3, 2},
	}

	for _, c := range tc {
		Sort(c, func(i, j int) bool {
			return c[i] < c[j]
		})

		n := len(c)
		for i := 1; i < n; i++ {
			assert.False(t, c[i-1] > c[i])
		}
	}
}
