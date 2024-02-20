package object

import "testing"

func TestStringHashKey(t *testing.T) {
	first1 := &String{Value: "Hello World"}
	first2 := &String{Value: "Hello World"}
	second1 := &String{Value: "My name is jonny 5"}
	second2 := &String{Value: "My name is jonny 5"}

	if first1.HashKey() != first2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if second1.HashKey() != second2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if first1.HashKey() == second2.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	first1 := &Integer{Value: 1}
	first2 := &Integer{Value: 1}
	second1 := &Integer{Value: 2}
	second2 := &Integer{Value: 2}

	if first1.HashKey() != first2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}
	if second1.HashKey() != second2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}
	if first1.HashKey() == second2.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	first1 := &Boolean{Value: true}
	first2 := &Boolean{Value: true}
	second1 := &Boolean{Value: false}
	second2 := &Boolean{Value: false}

	if first1.HashKey() != first2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}
	if second1.HashKey() != second2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}
	if first1.HashKey() == second2.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}
