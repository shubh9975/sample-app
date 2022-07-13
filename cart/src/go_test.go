package main

import (
	"testing"
        "cart/src/mypkg"
	//"test/mypkg"
)
func TestHello(t *testing.T) {
	got := mypkg.PrintHello()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
