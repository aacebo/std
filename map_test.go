package std

import "testing"

func TestMapSet(t *testing.T) {
	m := NewMap[string]()
	m.Set("key-1", "test")

	if *m.Get("key-1") != "test" {
		t.Fatal("key/value does not match expected")
	}
}

func TestMapHas(t *testing.T) {
	m := NewMap[string]()
	m.Set("key-1", "test")

	if !m.Has("key-1") {
		t.Fatal("map should have \"key-1\"")
	}

	if m.Has("key-2") {
		t.Fatal("map should not have \"key-2\"")
	}
}

func TestMapRemove(t *testing.T) {
	m := NewMap[string]()
	m.Set("key-1", "test")

	if !m.Has("key-1") {
		t.Fatal("map should have \"key-1\"")
	}

	m.Remove("key-1")

	if m.Has("key-1") {
		t.Fatal("map should not have \"key-1\" after remove")
	}
}
