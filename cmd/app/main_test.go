package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	var buf bytes.Buffer
	greet(&buf)

	if buf.String() != "Hello, GitHub Actions!" {
		t.Fail()
	}
	//t.Fail()
}
