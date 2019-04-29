package fullsort

import "testing"

func stringPanic(t *testing.T, a, b interface{}, msg string) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(msg + ": expected panic")
		}
	}
	stringComparator(a, b)
}

func TestStringComparator(t *testing.T) {
	a, b := "aaaa", "ba"
	if stringComparator(a, b) {
		t.Error("Compare \"aaaa\" and \"ba\": expected false")
	}
	if !stringComparator(b, a) {
		t.Error("Compare \"ba\" and \"aaaa\": expected true")
	}
	if stringComparator("", "") {
		t.Error("Compare empty strings: expected false")
	}
	if stringComparator("abcdef", "abcdef") {
		t.Error("Compare \"abcdef\" and \"abcdef\": expected false")
	}
	if !stringComparator("ptdr", "") {
		t.Error("Compare \"ptdr\" and empty string: expected true")
	}
	if stringComparator("hell", "hello") {
		t.Error("Compare \"hell\" and \"hello\": expected false")
	}

	stringPanic(t, 42, 35, "Compare integers")
	stringPanic(t, "hello", 35, "Compare string and integer")
	stringPanic(t, 42, "world", "Compare integer and string")
}

func TestFloatComparator32(t *testing.T) {
	var a, b float32

	a, b = 4.231, 42.5
	if floatComparator(a, b) {
		t.Error("Compare 4.231 and 42.5: expected false")
	}
	if !floatComparator(b, a) {
		t.Error("Compare 42.5 and 4.231: expected true")
	}
	a, b = 0, 42
	if floatComparator(a, b) {
		t.Error("Compare 0 and 42: expected false")
	}
	if !floatComparator(b, a) {
		t.Error("Compare 42 and 0: expected true")
	}
	a, b = -64, 12
	if floatComparator(a, b) {
		t.Error("Compare -64 and 12: expected false")
	}
	if !floatComparator(b, a) {
		t.Error("Compare 12 and -64: expected true")
	}

	if a, b = 0, 0; floatComparator(a, b) {
		t.Error("Compare 0 and 0: expected false")
	}
	if a, b = -35, -35; floatComparator(a, b) {
		t.Error("Compare -35 and -35: expected false")
	}
}

func TestFloatComparator64(t *testing.T) {
	var a, b float64
	
	a, b = 4.231, 42.5
	if floatComparator(a, b) {
		t.Error("Compare 4.231 and 42.5: expected false")
	}
	if !floatComparator(b, a) {
		t.Error("Compare 42.5 and 4.231: expected true")
	}
	a, b = 0, 42
	if floatComparator(a, b) {
		t.Error("Compare 0 and 42: expected false")
	}
	if !floatComparator(b, a) {
		t.Error("Compare 42 and 0: expected true")
	}
	a, b = -64, 12
	if floatComparator(a, b) {
		t.Error("Compare -64 and 12: expected false")
	}
	if !floatComparator(b, a) {
		t.Error("Compare 12 and -64: expected true")
	}

	if a, b = 0, 0; floatComparator(a, b) {
		t.Error("Compare 0 and 0: expected false")
	}
	if a, b = -35, -35; floatComparator(a, b) {
		t.Error("Compare -35 and -35: expected false")
	}
}