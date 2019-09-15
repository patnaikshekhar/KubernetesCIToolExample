package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	if result != 5 {
		t.Errorf("Error adding numbers")
	}
}
