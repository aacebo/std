package std

import "testing"

func TestNewSet(t *testing.T) {
	set := NewSet[int]()

	if set.Size() > 0 {
		t.Fatal("set should be empty when initialized")
	}
}

func TestSetAdd(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Fatal("set should be of size 3")
	}
}

func TestSetRemove(t *testing.T) {
	set := NewSet[string]()
	set.Add("one")
	set.Add("two")
	set.Add("two")
	set.Add("three")
	set.Remove(1)

	if set.Size() != 2 {
		t.Fatal("set should be of size 2")
	}
}

func TestSetAt(t *testing.T) {
	set := NewSet[string]()
	set.Add("one")
	set.Add("two")
	set.Add("two")
	set.Add("three")

	if *set.At(2) != "three" {
		t.Fatal("set value at index 2 should equal \"three\"")
	}
}
