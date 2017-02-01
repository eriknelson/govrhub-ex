package foo

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	value := Hello()
	fmt.Println("thing")
	if value != "Foo::Hello" {
		t.Fatal("Got unexpected value for Foo::Hello")
	}
}

// NOTE: World is untested.
