package main

import "testing"

func TestSomething(t *testing.T) {
	if add(2, 2) != 4 {
		t.Fatal("add(2, 2) should return 4")
	}
}
