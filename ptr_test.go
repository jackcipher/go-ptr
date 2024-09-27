package ptr

import (
	"testing"
	"time"
)

func TestPtr(t *testing.T) {
	v := 42
	p := Ptr(v)
	if *p != v {
		t.Errorf("Ptr(%d) = %d, want %d", v, *p, v)
	}
}

func TestString(t *testing.T) {
	v := "hello"
	p := String(v)
	if *p != v {
		t.Errorf("String(%q) = %q, want %q", v, *p, v)
	}
}

func TestInt(t *testing.T) {
	v := 42
	p := Int(v)
	if *p != v {
		t.Errorf("Int(%d) = %d, want %d", v, *p, v)
	}
}

func TestBool(t *testing.T) {
	v := true
	p := Bool(v)
	if *p != v {
		t.Errorf("Bool(%t) = %t, want %t", v, *p, v)
	}
}

func TestFloat64(t *testing.T) {
	v := 3.14
	p := Float64(v)
	if *p != v {
		t.Errorf("Float64(%f) = %f, want %f", v, *p, v)
	}
}

func TestTime(t *testing.T) {
	v := time.Now()
	p := Time(v)
	if *p != v {
		t.Errorf("Time(%v) = %v, want %v", v, *p, v)
	}
}

func TestSlicePtr(t *testing.T) {
	v := []int{1, 2, 3}
	p := SlicePtr(v)
	if len(*p) != len(v) {
		t.Errorf("SlicePtr(%v) = %v, want %v", v, *p, v)
	}
}

func TestMapPtr(t *testing.T) {
	v := map[string]int{"a": 1, "b": 2}
	p := MapPtr(v)
	if len(*p) != len(v) {
		t.Errorf("MapPtr(%v) = %v, want %v", v, *p, v)
	}
}

func TestIsNil(t *testing.T) {
	var p *int
	if !IsNil(p) {
		t.Errorf("IsNil(nil) = false, want true")
	}
	v := 42
	p = &v
	if IsNil(p) {
		t.Errorf("IsNil(non-nil) = true, want false")
	}
}

func TestGetValueOrDefault(t *testing.T) {
	v := 42
	p := &v
	if GetValueOrDefault(p, 0) != v {
		t.Errorf("GetValueOrDefault(%d, 0) = %d, want %d", v, GetValueOrDefault(p, 0), v)
	}
	var nilP *int
	if GetValueOrDefault(nilP, 0) != 0 {
		t.Errorf("GetValueOrDefault(nil, 0) = %d, want 0", GetValueOrDefault(nilP, 0))
	}
}

func TestGetNonZeroValueOrDefault(t *testing.T) {
	v := 42
	p := &v
	if GetNonZeroValueOrDefault(p, 0) != v {
		t.Errorf("GetNonZeroValueOrDefault(%d, 0) = %d, want %d", v, GetNonZeroValueOrDefault(p, 0), v)
	}
	zero := 0
	zeroP := &zero
	if GetNonZeroValueOrDefault(zeroP, 42) != 42 {
		t.Errorf("GetNonZeroValueOrDefault(0, 42) = %d, want 42", GetNonZeroValueOrDefault(zeroP, 42))
	}
}

func TestPtrIf(t *testing.T) {
	v := 42
	p := NewPtrIf(true, v)
	if p == nil || *p != v {
		t.Errorf("PtrIf(true, %d) = %v, want %d", v, p, v)
	}
	p = NewPtrIf(false, v)
	if p != nil {
		t.Errorf("PtrIf(false, %d) = %v, want nil", v, p)
	}
}

func TestEqual(t *testing.T) {
	a, b := 42, 42
	if !Equal(&a, &b) {
		t.Errorf("Equal(%d, %d) = false, want true", a, b)
	}
	c := 43
	if Equal(&a, &c) {
		t.Errorf("Equal(%d, %d) = true, want false", a, c)
	}
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{2, 4, 6}
	result := Map(input, func(v int) int { return v * 2 })
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Map()[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}
	result := Filter(input, func(v int) bool { return v%2 == 0 })
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Filter()[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

// TestMapKeys tests the MapKeys function.
func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := []string{"a", "b", "c"}
	keys := MapKeys(m)

	for _, key := range expected {
		found := false
		for _, k := range keys {
			if k == key {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected key %s not found", key)
		}
	}
}

// TestReduce tests the Reduce function.
func TestReduce(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	sum := Reduce(slice, func(acc int, v int) int {
		return acc + v
	}, 0)

	if sum != 10 {
		t.Errorf("Expected sum to be 10, got %d", sum)
	}
}

// TestFirstOrDefault tests the FirstOrDefault function.
func TestFirstOrDefault(t *testing.T) {
	slice := []int{1, 2, 3}
	defaultValue := 0
	result := FirstOrDefault(slice, defaultValue)

	if result != 1 {
		t.Errorf("Expected first element to be 1, got %d", result)
	}

	emptySlice := []int{}
	resultEmpty := FirstOrDefault(emptySlice, defaultValue)

	if resultEmpty != defaultValue {
		t.Errorf("Expected default value %d, got %d", defaultValue, resultEmpty)
	}
}

