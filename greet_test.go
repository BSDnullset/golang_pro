package main

import "testing"

func test_greet(t *testing.T) {
	got := greet()
	want := "gsb"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
