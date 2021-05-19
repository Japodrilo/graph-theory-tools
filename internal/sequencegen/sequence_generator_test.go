package sequencegen

import (
	"reflect"
	"testing"
)

func contains(s [][]bool, e []bool) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

func TestWeight(t *testing.T) {
	// Test 1: weight 0
	gen := Weight(3, 0)
	res := make([]bool, 3)
	got := gen()
	if !reflect.DeepEqual(res, got) {
		t.Errorf("Expected %v, got %v", res, got)
	}
	got = gen()
	if !reflect.DeepEqual(res, got) {
		t.Errorf("Expected %v, got %v", res, got)
	}
	// Test 2: weight 1
	got2 := make([][]bool, 3)
	gen = Weight(3, 1)
	for i := 0; i < 3; i++ {
		got2[i] = gen()
	}
	elem := []bool{true, false, false}
	found := contains(got2, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got2)
	}
	elem = []bool{false, true, false}
	found = contains(got2, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got2)
	}
	elem = []bool{false, false, true}
	found = contains(got2, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got2)
	}
	// Test 3: weight 2
	got3 := make([][]bool, 3)
	gen = Weight(3, 2)
	for i := 0; i < 3; i++ {
		got3[i] = gen()
	}
	elem = []bool{false, true, true}
	found = contains(got3, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got3)
	}
	elem = []bool{true, false, true}
	found = contains(got3, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got3)
	}
	elem = []bool{true, true, false}
	found = contains(got3, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got3)
	}
	// Test 4: Weight 3
	gen = Weight(3, 3)
	res = []bool{true, true, true}
	got = gen()
	if !reflect.DeepEqual(res, got) {
		t.Errorf("Expected %v, got %v", res, got)
	}
	got = gen()
	if !reflect.DeepEqual(res, got) {
		t.Errorf("Expected %v, got %v", res, got)
	}
}
