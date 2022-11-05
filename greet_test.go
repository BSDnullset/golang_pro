package main

import "testing"

// remember to run:$ go mod init SOMENAME
// in each new folder before running commands like: $ go test
// or: $ go build

func TestGreet(t *testing.T) {
	got := greet()
	want := "clusters slice segregation"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}

	get := displayID_number()
	wanted := "clusters have unique identifiers"

	if get != wanted {
		t.Errorf("got %q, wanted %q", get, wanted)
	}

}

// go test files need to be named someVariable_test.go
// and the Test function must start with the word Test so that it gets recognized by the linker & compiler
// the function TestSomeVariable takes ONLY one argument: t *testing.T
// for the compiler to recognize   *testing.T	one first needs to		import "testing"
// just like you'd import "fmt"
