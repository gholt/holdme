// Automatically generated with: gastly orderednumbersdups_test.go ../orderedintsdups_test.go holdme NumericType=droptype:int Number=Int number=int

package holdme

import "testing"

func Test_OrderedIntsDups_Add_inOrder(t *testing.T) {
	var slice OrderedIntsDups
	for v := int(0); v < 10; v++ {
		slice = slice.Add(v)
	}
	if slice.String() != "[0 1 2 3 4 5 6 7 8 9]" {
		t.Fatal(slice.String())
	}
}

func Test_OrderedIntsDups_Add_inReverseOrder(t *testing.T) {
	var slice OrderedIntsDups
	for v := int(9); ; v-- {
		slice = slice.Add(v)
		if v == 0 {
			break
		}
	}
	if slice.String() != "[0 1 2 3 4 5 6 7 8 9]" {
		t.Fatal(slice.String())
	}
}

func Test_OrderedIntsDups_Add_inShuffledOrder(t *testing.T) {
	var slice OrderedIntsDups
	slice = slice.Add(3)
	slice = slice.Add(6)
	slice = slice.Add(4)
	slice = slice.Add(9)
	slice = slice.Add(1)
	slice = slice.Add(5)
	slice = slice.Add(2)
	slice = slice.Add(7)
	slice = slice.Add(0)
	slice = slice.Add(8)
	if slice.String() != "[0 1 2 3 4 5 6 7 8 9]" {
		t.Fatal(slice.String())
	}
}

func Test_OrderedIntsDups_Add_inShuffledOrderWithDups(t *testing.T) {
	var slice OrderedIntsDups
	slice = slice.Add(3)
	slice = slice.Add(6)
	slice = slice.Add(4)
	slice = slice.Add(9)
	slice = slice.Add(1)
	slice = slice.Add(5)
	slice = slice.Add(2)
	slice = slice.Add(7)
	slice = slice.Add(0)
	slice = slice.Add(8)
	slice = slice.Add(4)
	slice = slice.Add(7)
	slice = slice.Add(6)
	slice = slice.Add(1)
	slice = slice.Add(2)
	slice = slice.Add(9)
	slice = slice.Add(8)
	slice = slice.Add(5)
	slice = slice.Add(3)
	slice = slice.Add(0)
	slice = slice.Add(8)
	slice = slice.Add(8)
	slice = slice.Add(8)
	slice = slice.Add(8)
	slice = slice.Add(8)
	if slice.String() != "[0 0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 8 8 8 8 8 9 9]" {
		t.Fatal(slice.String())
	}
}

func Test_OrderedIntsDups_Remove(t *testing.T) {
	var slice OrderedIntsDups
	for v := int(0); v < 10; v++ {
		slice = slice.Add(v)
	}
	if slice.String() != "[0 1 2 3 4 5 6 7 8 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(0)
	if slice.String() != "[1 2 3 4 5 6 7 8 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(5)
	if slice.String() != "[1 2 3 4 6 7 8 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(9)
	if slice.String() != "[1 2 3 4 6 7 8]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(1000)
	if slice.String() != "[1 2 3 4 6 7 8]" {
		t.Fatal(slice.String())
	}
}

func Test_OrderedIntsDups_Remove_hadDups(t *testing.T) {
	var slice OrderedIntsDups
	for v := int(0); v < 10; v++ {
		slice = slice.Add(v)
	}
	for v := int(0); v < 10; v++ {
		slice = slice.Add(v)
	}
	if slice.String() != "[0 0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 9 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(0)
	if slice.String() != "[0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 9 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(5)
	if slice.String() != "[0 1 1 2 2 3 3 4 4 5 6 6 7 7 8 8 9 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(9)
	if slice.String() != "[0 1 1 2 2 3 3 4 4 5 6 6 7 7 8 8 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(1000)
	if slice.String() != "[0 1 1 2 2 3 3 4 4 5 6 6 7 7 8 8 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(0)
	if slice.String() != "[1 1 2 2 3 3 4 4 5 6 6 7 7 8 8 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(5)
	if slice.String() != "[1 1 2 2 3 3 4 4 6 6 7 7 8 8 9]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(9)
	if slice.String() != "[1 1 2 2 3 3 4 4 6 6 7 7 8 8]" {
		t.Fatal(slice.String())
	}
	slice = slice.Remove(1000)
	if slice.String() != "[1 1 2 2 3 3 4 4 6 6 7 7 8 8]" {
		t.Fatal(slice.String())
	}
}
