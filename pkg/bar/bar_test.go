package bar

import (
	"testing"
)

func TestBar(t *testing.T) {
	value := Hello()
	if value != "Bar::Hello" {
		t.Fatal("Got unexpected value for Bar::Hello")
	}
}
