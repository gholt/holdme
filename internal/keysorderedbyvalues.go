package internal

import "fmt"

type NumericKeyType int
type NumericValueType int

// KeysOrderedByValues keeps Keys in order by Values[Keys[i]] value.
//
// It assumes a relatively static set in Keys that you want to keep sorted as
// their Values change. You first add all the keys using Add and then use Move
// to change the values. As you do so, Keys will be kept in order.
//
// Note that the values of the keys are simply stored in a slice and not a map,
// so if you have a high key, it will use a lot of memory. It should be easy to
// modify to use a map if desired though; for my use, a slice was fine and much
// faster than a map.
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
type KeysOrderedByValues struct {
	Keys     []NumericKeyType
	Values   []NumericValueType
	RandIntn func(int) int
}

func NewKeysOrderedByValues(cap int, randIntn func(int) int) *KeysOrderedByValues {
	return &KeysOrderedByValues{
		Keys:     make([]NumericKeyType, 0, cap),
		Values:   make([]NumericValueType, 0, cap),
		RandIntn: randIntn,
	}
}

func (x *KeysOrderedByValues) Add(key NumericKeyType, value NumericValueType) {
	ln := len(x.Values)
	if int(key) < ln {
		x.Move(key, value)
		return
	}
	if int(key) == ln {
		x.Values = append(x.Values, value)
	} else {
		x.Values = append(x.Values, make([]NumericValueType, int(key)-ln+1)...)
		x.Values[key] = value
	}
	right := 0
	hi := len(x.Keys)
	for right < hi {
		mid := (right + hi) / 2
		if value > x.Values[x.Keys[mid]] {
			hi = mid
		} else {
			right = mid + 1
		}
	}
	if x.RandIntn != nil {
		left := 0
		hi = len(x.Keys)
		for left < hi {
			mid := (left + hi) / 2
			if x.Values[x.Keys[mid]] > value {
				left = mid + 1
			} else {
				hi = mid
			}
		}
		if right-left > 2 {
			right = right - x.RandIntn(right-left)
		}
	}
	x.Keys = append(x.Keys, 0)
	copy(x.Keys[right+1:], x.Keys[right:])
	x.Keys[right] = key
}

func (x *KeysOrderedByValues) Move(key NumericKeyType, value NumericValueType) {
	var oldPosition int
	oldValue := x.Values[key]
	right := 0
	hi := len(x.Keys)
	for right < hi {
		mid := (right + hi) / 2
		if oldValue > x.Values[x.Keys[mid]] {
			hi = mid
		} else {
			right = mid + 1
		}
	}
	right--
	if x.Keys[right] == key {
		oldPosition = right
	} else {
		left := 0
		hi = len(x.Keys)
		for left < hi {
			mid := (left + hi) / 2
			if x.Values[x.Keys[mid]] > oldValue {
				left = mid + 1
			} else {
				hi = mid
			}
		}
		if x.Keys[left] == key {
			oldPosition = left
		} else {
			for oldPosition = left + 1; x.Keys[oldPosition] != key && oldPosition < right; oldPosition++ {
			}
		}
	}
	right = 0
	hi = len(x.Keys)
	for right < hi {
		mid := (right + hi) / 2
		if value > x.Values[x.Keys[mid]] {
			hi = mid
		} else {
			right = mid + 1
		}
	}
	if x.RandIntn != nil {
		left := 0
		hi = len(x.Keys)
		for left < hi {
			mid := (left + hi) / 2
			if x.Values[x.Keys[mid]] > value {
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
		copy(x.Keys[right+1:], x.Keys[right:oldPosition])
		x.Keys[right] = key
	} else if oldPosition < right {
		copy(x.Keys[oldPosition:], x.Keys[oldPosition+1:right])
		x.Keys[right-1] = key
	}
	x.Values[key] = value
}

func (x *KeysOrderedByValues) String() string {
	s := "["
	for i, key := range x.Keys {
		if i == 0 {
			s += fmt.Sprintf("%d:%d", key, x.Values[key])
		} else {
			s += fmt.Sprintf(" %d:%d", key, x.Values[key])
		}
	}
	return s + "]"
}
