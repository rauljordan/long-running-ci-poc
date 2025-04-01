package main

import "testing"

func TestSomething(t *testing.T) {
	if add(2, 2) != 4 {
		t.Fatal("add(2, 2) should return 4")
	}
}

func TestOther(t *testing.T) {
	if add(1, 1) != 2 {
		t.Fatal("add(1, 1) should return 2")
	}
}
