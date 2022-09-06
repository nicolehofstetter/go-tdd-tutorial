package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tommy")
	want := "Hello Tommy!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
