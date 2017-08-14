package holdme

import (
	"math/rand"
	"testing"
)

func helperIntValueOrderedIntKeysInOrder(t *testing.T, x *IntValueOrderedIntKeys) {
	f := true
	var pk int
	var pv int
	for i, k := range x.OrderedKeys {
		v := x.KeyToValue[k]
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

func helperIntValueOrderedIntKeysFind(t *testing.T, x *IntValueOrderedIntKeys, k int) int {
	for i, k2 := range x.OrderedKeys {
		if k2 == k {
			return i
		}
	}
	t.Fatalf("could not find %v", k)
	return -1
}

func Test_IntValueOrderedIntKeys_Add(t *testing.T) {
	var x *IntValueOrderedIntKeys
	refresh := func() {
		x = &IntValueOrderedIntKeys{}
		x.Add(0, 10)
		for k := int(1); k < 6; k++ {
			x.Add(k, 0)
		}
		x.Add(6, -10)
		helperIntValueOrderedIntKeysFind(t, x, 0)
		for k := int(1); k < 6; k++ {
			helperIntValueOrderedIntKeysFind(t, x, k)
		}
		helperIntValueOrderedIntKeysFind(t, x, 6)
	}
	refresh()
	x.RandIntn = nil
	x.Add(7, 0)
	helperIntValueOrderedIntKeysInOrder(t, x)
	i := helperIntValueOrderedIntKeysFind(t, x, 7)
	randIntn := rand.New(rand.NewSource(0)).Intn
	i2 := i
	for j := 0; j < 100; j++ {
		refresh()
		x.RandIntn = randIntn
		x.Add(7, 0)
		helperIntValueOrderedIntKeysInOrder(t, x)
		i2 = helperIntValueOrderedIntKeysFind(t, x, 7)
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
		helperIntValueOrderedIntKeysInOrder(t, x)
		i3 = helperIntValueOrderedIntKeysFind(t, x, 7)
		if i3 != i && i3 != i2 {
			break
		}
	}
	if i3 == i || i3 == i2 {
		t.Fatal("expected positions to change", x)
	}
	refresh()
	x.Add(7, 0)
	helperIntValueOrderedIntKeysInOrder(t, x)
	ln := len(x.OrderedKeys)
	x.Add(7, 0)
	helperIntValueOrderedIntKeysInOrder(t, x)
	ln2 := len(x.OrderedKeys)
	if ln != ln2 {
		t.Fatal("length changed", ln, ln2)
	}
}

func Test_IntValueOrderedIntKeys_Move(t *testing.T) {
	x := &IntValueOrderedIntKeys{}
	for k := int(1); k < 10; k++ {
		x.Add(k, int(k))
	}
	helperIntValueOrderedIntKeysInOrder(t, x)
	x.Move(3, 7)
	helperIntValueOrderedIntKeysInOrder(t, x)
	if x.KeyToValue[3] != 7 {
		t.Fatal("value was", x.KeyToValue[3])
	}
	x.Move(3, 7)
	helperIntValueOrderedIntKeysInOrder(t, x)
	if x.KeyToValue[3] != 7 {
		t.Fatal("value was", x.KeyToValue[3])
	}
	for k := int(10); k < 15; k++ {
		x.Add(k, 5)
	}
	helperIntValueOrderedIntKeysInOrder(t, x)
	x.Move(7, 3)
	helperIntValueOrderedIntKeysInOrder(t, x)
	x.Move(7, 5)
	helperIntValueOrderedIntKeysInOrder(t, x)
	i := helperIntValueOrderedIntKeysFind(t, x, 7)
	x.RandIntn = rand.New(rand.NewSource(0)).Intn
	i2 := i
	for j := 0; j < 100; j++ {
		x.Move(7, 3)
		helperIntValueOrderedIntKeysInOrder(t, x)
		x.Move(7, 5)
		helperIntValueOrderedIntKeysInOrder(t, x)
		i2 = helperIntValueOrderedIntKeysFind(t, x, 7)
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
		helperIntValueOrderedIntKeysInOrder(t, x)
		x.Move(7, 5)
		helperIntValueOrderedIntKeysInOrder(t, x)
		i3 = helperIntValueOrderedIntKeysFind(t, x, 7)
		if i3 != i && i3 != i2 {
			break
		}
	}
	if i3 == i || i3 == i2 {
		t.Fatal("expected positions to change", x)
	}
}
