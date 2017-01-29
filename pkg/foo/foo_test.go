package foo

import (
	"testing"
)

func TestFoo(t *testing.T) {
	value := Hello()
	if value != "Foo::Hello" {
		t.Fatal("Got unexpected value for Foo::Hello")
	}
}

// NOTE: World is untested.
