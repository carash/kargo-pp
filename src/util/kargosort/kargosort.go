package kargosort

import (
	"reflect"
)

func qpart(arr reflect.Value, less func(i, j int) bool, low, high int) int {
	i := low
	j := high

	for {
		// Find leftmost element greater than
		// or equal to pivot

		for less(i, low) {
			i++
		}

		// Find rightmost element smaller than
		// or equal to pivot
		for less(low, j) {
			j--
		}

		// If two pointers met.
		if i >= j {
			return j
		}

		templow := arr.Index(i).Interface()
		temphigh := arr.Index(j).Interface()

		arr.Index(i).Set(reflect.ValueOf(temphigh))
		arr.Index(j).Set(reflect.ValueOf(templow))

		i++
		j--
	}
}

// arr is value containing slice
func qsort(arr reflect.Value, less func(i, j int) bool, low, high int) {
	if low < high {
		pi := qpart(arr, less, low, high)

		qsort(arr, less, low, pi)
		qsort(arr, less, pi+1, high)
	}
}

func Sort(arr interface{}, less func(i, j int) bool) {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		qsort(s, less, 0, s.Len()-1)

	default:
		return
	}
}
