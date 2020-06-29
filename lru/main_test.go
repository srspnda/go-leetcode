package main

import "testing"

func TestLRUCache(t *testing.T) {
	l := NewLRUCache(2)

	l.Put(1, 1)
	l.Put(2, 2)
	if l.Get(1) != 1 {
		t.Fatalf("cache get 1 should equal 1")
	}
	l.Put(3, 3)
	if l.Get(2) != -1 {
		t.Fatalf("get 2 should be -1 since cache is {1, 3}")
	}
	l.Put(4, 4)
	if l.Get(1) != -1 {
		t.Fatalf("get 1 should be -1 since cache is {3, 4}")
	}
	if l.Get(3) != 3 {
		t.Fatalf("get 3 should return 3")
	}
	if l.Get(4) != 4 {
		t.Fatalf("get 4 should return 4")
	}
}
