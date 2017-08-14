package holdme

import "fmt"

// OrderedUint32sDups keeps a slice of uint32s in order, allowing duplicate values.
type OrderedUint32sDups []uint32

func (slice OrderedUint32sDups) Add(v uint32) OrderedUint32sDups {
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
	if left == ln {
		return append(slice, v)
	}
	slice = append(slice, 0)
	copy(slice[left+1:], slice[left:])
	slice[left] = v
	return slice
}

func (slice OrderedUint32sDups) Remove(v uint32) OrderedUint32sDups {
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

func (slice OrderedUint32sDups) String() string {
	return fmt.Sprintf("%v", []uint32(slice))
}
