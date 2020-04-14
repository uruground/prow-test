package main

import "testing"

func TestEcho(t *testing.T) {
	s := "hi"
	if s != echo(s) {
		t.Fatalf("echo not match")
	}
}
