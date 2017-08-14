package holdme

import "fmt"

// OrderedUint32sNoDups keeps a slice of uint32s in order, without duplicates.
type OrderedUint32sNoDups []uint32

func (slice OrderedUint32sNoDups) Add(v uint32) OrderedUint32sNoDups {
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

func (slice OrderedUint32sNoDups) Remove(v uint32) OrderedUint32sNoDups {
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

func (slice OrderedUint32sNoDups) String() string {
	return fmt.Sprintf("%v", []uint32(slice))
}
