package main

import "testing"

// remember to run:$ go mod init SOMENAME
// in each new folder before running commands like: $ go test
// or: $ go build

func TestGreet(t *testing.T) {
	got := greet()
	want := "clusters slice segregation"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
