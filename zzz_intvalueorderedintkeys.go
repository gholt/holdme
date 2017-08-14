package holdme

import "fmt"

// IntValueOrderedIntKeys keeps OrderedKeys in order by KeyToValue[OrderedKeys[i]] value.
//
// It assumes a relatively static set in OrderedKeys that you want to keep sorted as
// their values in KeyToValue change. You first add all the keys using Add and then
// use Move to change the values. As you do so, OrderedKeys will be kept in order.
//
// Note that the values of the keys are simply stored in a slice and not a map,
// so wildly varying min/max keys will use a lot of memory. It should be easy
// to modify to use a map if desired though; for my use, a slice was fine and
// much faster than a map.
//
// The RandIntn function may be set to randomize key locations whose values are
// the same. Even if RandIntn is nil, the order of keys whose values are the
// same is not guaranteed to be in any particular order (think like Go's map
// key order, no guarantees).
//
// I know it's a weird "key/value randomization bag" or something and probably
// not generally useful, but I had to create it for my
// github.com/gholt/ring/lowring package where I needed to sort nodes and
// groups by their desires, so here it is.
//
// This code will copy memory quite a bit and can probably be improved with
// scanning for value runs and swapping. Of course, if there are few runs,
// swapping "by hand" will probably be slower than just calling copy. Not sure
// yet, just haven't gotten that far. Oh, and another note if you (or future
// me) decides to try this optimization, you'll have to swap before applying
// the RandIntn, and then use that to swap a second time.
type IntValueOrderedIntKeys struct {
	OrderedKeys []int
	KeyToValue  []int
	RandIntn    func(int) int
}

func NewIntValueOrderedIntKeys(cap int, randIntn func(int) int) *IntValueOrderedIntKeys {
	return &IntValueOrderedIntKeys{
		OrderedKeys: make([]int, 0, cap),
		KeyToValue:  make([]int, 0, cap),
		RandIntn:    randIntn,
	}
}

func (x *IntValueOrderedIntKeys) Add(key int, value int) {
	ln := len(x.KeyToValue)
	if int(key) < ln {
		x.Move(key, value)
		return
	}
	if int(key) == ln {
		x.KeyToValue = append(x.KeyToValue, value)
	} else {
		x.KeyToValue = append(x.KeyToValue, make([]int, int(key)-ln+1)...)
		x.KeyToValue[key] = value
	}
	right := 0
	hi := len(x.OrderedKeys)
	for right < hi {
		mid := (right + hi) / 2
		if value > x.KeyToValue[x.OrderedKeys[mid]] {
			hi = mid
		} else {
			right = mid + 1
		}
	}
	if x.RandIntn != nil {
		left := 0
		hi = len(x.OrderedKeys)
		for left < hi {
			mid := (left + hi) / 2
			if x.KeyToValue[x.OrderedKeys[mid]] > value {
				left = mid + 1
			} else {
				hi = mid
			}
		}
		if right-left > 2 {
			right = right - x.RandIntn(right-left)
		}
	}
	x.OrderedKeys = append(x.OrderedKeys, 0)
	copy(x.OrderedKeys[right+1:], x.OrderedKeys[right:])
	x.OrderedKeys[right] = key
}

func (x *IntValueOrderedIntKeys) Move(key int, value int) {
	var oldPosition int
	oldValue := x.KeyToValue[key]
	right := 0
	hi := len(x.OrderedKeys)
	for right < hi {
		mid := (right + hi) / 2
		if oldValue > x.KeyToValue[x.OrderedKeys[mid]] {
			hi = mid
		} else {
			right = mid + 1
		}
	}
	right--
	if x.OrderedKeys[right] == key {
		oldPosition = right
	} else {
		left := 0
		hi = len(x.OrderedKeys)
		for left < hi {
			mid := (left + hi) / 2
			if x.KeyToValue[x.OrderedKeys[mid]] > oldValue {
				left = mid + 1
			} else {
				hi = mid
			}
		}
		if x.OrderedKeys[left] == key {
			oldPosition = left
		} else {
			for oldPosition = left + 1; x.OrderedKeys[oldPosition] != key && oldPosition < right; oldPosition++ {
			}
		}
	}
	right = 0
	hi = len(x.OrderedKeys)
	for right < hi {
		mid := (right + hi) / 2
		if value > x.KeyToValue[x.OrderedKeys[mid]] {
			hi = mid
		} else {
			right = mid + 1
		}
	}
	if x.RandIntn != nil {
		left := 0
		hi = len(x.OrderedKeys)
		for left < hi {
			mid := (left + hi) / 2
			if x.KeyToValue[x.OrderedKeys[mid]] > value {
				left = mid + 1
			} else {
				hi = mid
			}
		}
		if right-left > 2 {
			right = right - x.RandIntn(right-left)
		}
	}
	if right < oldPosition {
		copy(x.OrderedKeys[right+1:], x.OrderedKeys[right:oldPosition])
		x.OrderedKeys[right] = key
	} else if oldPosition < right {
		copy(x.OrderedKeys[oldPosition:], x.OrderedKeys[oldPosition+1:right])
		x.OrderedKeys[right-1] = key
	}
	x.KeyToValue[key] = value
}

func (x *IntValueOrderedIntKeys) String() string {
	s := "["
	for i, key := range x.OrderedKeys {
		if i == 0 {
			s += fmt.Sprintf("%d:%d", key, x.KeyToValue[key])
		} else {
			s += fmt.Sprintf(" %d:%d", key, x.KeyToValue[key])
		}
	}
	return s + "]"
}
