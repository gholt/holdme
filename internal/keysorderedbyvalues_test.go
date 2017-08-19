package internal

import (
	"math/rand"
	"testing"
)

func helperKeysOrderedByValuesInOrder(t *testing.T, x *KeysOrderedByValues) {
	f := true
	var pk NumericKeyType
	var pv NumericValueType
	for i, k := range x.Keys {
		v := x.Values[k]
		if f {
			f = false
		} else {
			if v > pv {
				t.Fatalf("%v:%v at index %d was greater than %v:%v at index %d", k, v, i, pk, pv, i-1)
			}
		}
		pk = k
		pv = v
	}
}

func helperKeysOrderedByValuesFind(t *testing.T, x *KeysOrderedByValues, k NumericKeyType) int {
	for i, k2 := range x.Keys {
		if k2 == k {
			return i
		}
	}
	t.Fatalf("could not find %v", k)
	return -1
}

func Test_KeysOrderedByValues_Add(t *testing.T) {
	var x *KeysOrderedByValues
	refresh := func() {
		x = &KeysOrderedByValues{}
		x.Add(0, 10)
		for k := NumericKeyType(1); k < 6; k++ {
			x.Add(k, 0)
		}
		x.Add(6, -10)
		helperKeysOrderedByValuesFind(t, x, 0)
		for k := NumericKeyType(1); k < 6; k++ {
			helperKeysOrderedByValuesFind(t, x, k)
		}
		helperKeysOrderedByValuesFind(t, x, 6)
	}
	refresh()
	x.RandIntn = nil
	x.Add(7, 0)
	helperKeysOrderedByValuesInOrder(t, x)
	i := helperKeysOrderedByValuesFind(t, x, 7)
	randIntn := rand.New(rand.NewSource(0)).Intn
	i2 := i
	for j := 0; j < 100; j++ {
		refresh()
		x.RandIntn = randIntn
		x.Add(7, 0)
		helperKeysOrderedByValuesInOrder(t, x)
		i2 = helperKeysOrderedByValuesFind(t, x, 7)
		if i2 != i {
			break
		}
	}
	if i2 == i {
		t.Fatal("expected positions to change", x)
	}
	randIntn = rand.New(rand.NewSource(1)).Intn
	i3 := i
	for j := 0; j < 100; j++ {
		refresh()
		x.RandIntn = randIntn
		x.Add(7, 0)
		helperKeysOrderedByValuesInOrder(t, x)
		i3 = helperKeysOrderedByValuesFind(t, x, 7)
		if i3 != i && i3 != i2 {
			break
		}
	}
	if i3 == i || i3 == i2 {
		t.Fatal("expected positions to change", x)
	}
	refresh()
	x.Add(7, 0)
	helperKeysOrderedByValuesInOrder(t, x)
	ln := len(x.Keys)
	x.Add(7, 0)
	helperKeysOrderedByValuesInOrder(t, x)
	ln2 := len(x.Keys)
	if ln != ln2 {
		t.Fatal("length changed", ln, ln2)
	}
}

func Test_KeysOrderedByValues_Move(t *testing.T) {
	x := &KeysOrderedByValues{}
	for k := NumericKeyType(1); k < 10; k++ {
		x.Add(k, NumericValueType(k))
	}
	helperKeysOrderedByValuesInOrder(t, x)
	x.Move(3, 7)
	helperKeysOrderedByValuesInOrder(t, x)
	if x.Values[3] != 7 {
		t.Fatal("value was", x.Values[3])
	}
	x.Move(3, 7)
	helperKeysOrderedByValuesInOrder(t, x)
	if x.Values[3] != 7 {
		t.Fatal("value was", x.Values[3])
	}
	for k := NumericKeyType(10); k < 15; k++ {
		x.Add(k, 5)
	}
	helperKeysOrderedByValuesInOrder(t, x)
	x.Move(7, 3)
	helperKeysOrderedByValuesInOrder(t, x)
	x.Move(7, 5)
	helperKeysOrderedByValuesInOrder(t, x)
	i := helperKeysOrderedByValuesFind(t, x, 7)
	x.RandIntn = rand.New(rand.NewSource(0)).Intn
	i2 := i
	for j := 0; j < 100; j++ {
		x.Move(7, 3)
		helperKeysOrderedByValuesInOrder(t, x)
		x.Move(7, 5)
		helperKeysOrderedByValuesInOrder(t, x)
		i2 = helperKeysOrderedByValuesFind(t, x, 7)
		if i2 != i {
			break
		}
	}
	if i2 == i {
		t.Fatal("expected positions to change", x)
	}
	x.RandIntn = rand.New(rand.NewSource(1)).Intn
	i3 := i
	for j := 0; j < 100; j++ {
		x.Move(7, 3)
		helperKeysOrderedByValuesInOrder(t, x)
		x.Move(7, 5)
		helperKeysOrderedByValuesInOrder(t, x)
		i3 = helperKeysOrderedByValuesFind(t, x, 7)
		if i3 != i && i3 != i2 {
			break
		}
	}
	if i3 == i || i3 == i2 {
		t.Fatal("expected positions to change", x)
	}
	x.RandIntn = rand.New(rand.NewSource(2)).Intn
	x.Move(7, 5)
	i = helperKeysOrderedByValuesFind(t, x, 7)
	i2 = i
	for j := 0; j < 100; j++ {
		x.Move(7, 5)
		helperKeysOrderedByValuesInOrder(t, x)
		i2 = helperKeysOrderedByValuesFind(t, x, 7)
		if i2 != i {
			break
		}
	}
	if i2 == i {
		t.Fatal("expected positions to change", x)
	}
}
