package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {
	b := bytes.NewBufferString("w1 w2 w3 w4")
	e := 4                      // Expected
	r := count(b, false, false) // Result
	if e != r {
		t.Errorf("Expected: %d, got %d instead.\n", e, r)
	}
}

func TestLineCount(t *testing.T) {
	b := bytes.NewBufferString("l1\nl2\nl3")
	e := 3                     // Expected
	r := count(b, true, false) // Result
	if e != r {
		t.Errorf("Expected: %d, got %d instead.\n", e, r)
	}
}

func TestByteCount(t *testing.T) {
	b := bytes.NewBufferString("b1\nb2\nb3 b4 b5")
	e := 14                     // Expected
	r := count(b, false, true)  // Result
	if e != r {
		t.Errorf("Expected: %d, got %d instead.\n", e, r)
	}
}
