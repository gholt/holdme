package internal

import "fmt"

// OrderedNumbersNoDups keeps a list of numbers in order, without duplicates.
type OrderedNumbersNoDups []NumericType

func (slice OrderedNumbersNoDups) Add(v NumericType) OrderedNumbersNoDups {
	left := 0
	ln := len(slice)
	hi := ln
	for left < hi {
		mid := (left + hi) / 2
		if slice[mid] < v {
			left = mid + 1
		} else {
			hi = mid
		}
	}
	if left < ln && slice[left] == v {
		return slice
	}
	if left == ln {
		return append(slice, v)
	}
	slice = append(slice, 0)
	copy(slice[left+1:], slice[left:])
	slice[left] = v
	return slice
}

func (slice OrderedNumbersNoDups) Remove(v NumericType) OrderedNumbersNoDups {
	left := 0
	ln := len(slice)
	hi := ln
	for left < hi {
		mid := (left + hi) / 2
		if slice[mid] < v {
			left = mid + 1
		} else {
			hi = mid
		}
	}
	if left < ln && slice[left] == v {
		return append(slice[:left], slice[left+1:]...)
	}
	return slice
}

func (slice OrderedNumbersNoDups) String() string {
	return fmt.Sprintf("%v", []NumericType(slice))
}
