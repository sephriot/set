package set

import (
	"testing"
)

func inOrder(data []interface{}) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i].(int) >= data[i+1].(int) {
			return false
		}
	}

	return true
}

func TestSet_Insert(t *testing.T) {
	cmp := func(a, b interface{}) bool { return a.(int) < b.(int) }
	s := New(cmp)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	if !inOrder(s.data) {
		t.Fail()
	}

	s = New(cmp)
	s.Insert(4)
	s.Insert(3)
	s.Insert(2)
	s.Insert(1)
	if !inOrder(s.data) {
		t.Fail()
	}

	s = New(cmp)
	s.Insert(2)
	s.Insert(3)
	s.Insert(1)
	s.Insert(4)
	if !inOrder(s.data) {
		t.Fail()
	}

	s = New(cmp)
	s.Insert(3)
	s.Insert(1)
	s.Insert(2)
	s.Insert(4)
	if !inOrder(s.data) {
		t.Fail()
	}
}

func TestSet_Exists(t *testing.T) {
	cmp := func(a, b interface{}) bool { return a.(float64) < b.(float64) }
	s := New(cmp)
	s.Insert(3.0)
	s.Insert(1.0)
	s.Insert(2.0)
	s.Insert(4.0)

	if !s.Exists(1.0) {
		t.Log("Failed to find", 1.0)
		t.Fail()
	}

	if !s.Exists(2.0) {
		t.Log("Failed to find", 2.0)
		t.Fail()
	}

	if !s.Exists(3.0) {
		t.Log("Failed to find", 3.0)
		t.Fail()
	}

	if !s.Exists(4.0) {
		t.Log("Failed to find", 4.0)
		t.Fail()
	}

	if s.Exists(5.0) {
		t.Log("Failed to not find", 5.0)
		t.Fail()
	}
}

func TestSet_Erase(t *testing.T) {
	cmp := func(a, b interface{}) bool { return a.(int) < b.(int) }
	s := New(cmp)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.Erase(1)
	if !inOrder(s.data) {
		t.Fail()
	}

	s = New(cmp)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.Erase(2)
	if !inOrder(s.data) {
		t.Fail()
	}

	s = New(cmp)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.Erase(3)
	if !inOrder(s.data) {
		t.Fail()
	}

	s = New(cmp)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.Erase(4)
	if !inOrder(s.data) {
		t.Fail()
	}

	s = New(cmp)
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.Erase(1)
	s.Erase(4)
	s.Erase(3)
	s.Erase(2)
	if !s.Empty() {
		t.Fail()
	}
}
