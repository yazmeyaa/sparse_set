package sparse_set_test

import (
	"testing"

	"github.com/yazmeyaa/sparse_set"
)

func TestSparseSetAdd(t *testing.T) {
	ss := sparse_set.NewSparseSet(32)
	ss.Add(12)
	if ss.Contains(12) == false {
		t.Error("Value not added")
	}
}

func TestSparseSetContains(t *testing.T) {
	ss := sparse_set.NewSparseSet(32)
	ss.Add(16)

	notExisting := ss.Contains(10)
	existing := ss.Contains(16)
	defaultValue := ss.Contains(0)

	if notExisting == true {
		t.Error("Element must not be in sparse set")
		return
	}
	if existing == false {
		t.Error("Element must be in sparse set")
	}
	if defaultValue == true {
		t.Error("Element 0 must not be in sparse set")
	}
}

func TestSparseSetRange(t *testing.T) {
	ss := sparse_set.NewSparseSet(32)
	for idx := 0; idx < 10; idx++ {
		ss.Add(idx * 2)
	}

	expectedValues := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}

	i := 0
	ss.Range(func(id sparse_set.Id) {
		if expectedValues[i] != id {
			t.Errorf("Expected value: <%d>, recieved: <%d>", expectedValues[i], id)
		}
		i++
	})
}

func TestRemove(t *testing.T) {
	ss := sparse_set.NewSparseSet(32)

	ss.Add(12)
	exist := ss.Contains(12)
	if !exist {
		t.Error("Added value is not exist")
	}

	ss.Remove(12)
	exist = ss.Contains(12)
	if exist {
		t.Error("Value exist after element remove")
	}
}

func TestGetValues(t *testing.T) {
	/*
		---------- Test accessibility ----------
	*/
	ss := sparse_set.NewSparseSet(32)
	ss.Add(1)
	ss.Add(4)
	ss.Add(13)
	ss.Remove(4)
	ss.Add(10)

	expectedValues := []int{1, 13, 10}

	for idx, value := range ss.GetAll() {
		if expectedValues[idx] != value {
			t.Error("Wrong value found in slice")
		}
	}

	/*
		---------- Mutation test ----------
	*/

	slice := ss.GetAll()
	slice[0] = 14

	if !ss.Contains(1) {
		t.Error("Value must be in sparse set")
	}
}

func TestClear(t *testing.T) {
	ss := sparse_set.NewSparseSet(32)
	for i := 0; i < 32; i++ {
		ss.Add(i)
	}

	ss.Clear()
	if len(ss.GetAll()) != 0 {
		t.Error("Values in sparse set after flush (.Clear())")
	}
}
