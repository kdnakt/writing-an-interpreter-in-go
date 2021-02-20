package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	asc1 := &Integer{Value: 12345}
	asc2 := &Integer{Value: 12345}
	desc1 := &Integer{Value: 54321}
	desc2 := &Integer{Value: 54321}

	if asc1.HashKey() != asc2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if desc1.HashKey() != desc2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if asc1.HashKey() == desc1.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}
