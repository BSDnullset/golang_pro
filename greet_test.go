package main

import "testing"

func test_greet(t *testing.T) {
	got := greet()
	want := "clusters slice segregation" 

	if got != want {
		t.Errorf("got %q want %q, got, want)
	}
}
