// Automatically generated with: gastly orderednumbersnodups.go ../orderedintsnodups.go holdme NumericType=droptype:int Number=Int number=int

package holdme

import "fmt"

// OrderedIntsNoDups keeps a list of ints in order, without duplicates.
type OrderedIntsNoDups []int

func (slice OrderedIntsNoDups) Add(v int) OrderedIntsNoDups {
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

func (slice OrderedIntsNoDups) Remove(v int) OrderedIntsNoDups {
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

func (slice OrderedIntsNoDups) String() string {
	return fmt.Sprintf("%v", []int(slice))
}
