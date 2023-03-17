package main

import "testing"

func TestUnpack001(t *testing.T) {

	var got, want, ps string

	ps = "a4bc2d5e"
	want = "aaaabccddddde"
	got = StringUnpacking(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack002(t *testing.T) {

	var got, want, ps string

	ps = "abcd"
	want = "abcd"
	got = StringUnpacking(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack003(t *testing.T) {

	var got, want, ps string

	ps = "45"
	want = ""
	got = StringUnpacking(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack004(t *testing.T) {

	var got, want, ps string

	ps = ""
	want = ""
	got = StringUnpacking(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
