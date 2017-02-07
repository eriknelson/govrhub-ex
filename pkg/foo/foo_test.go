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

func TestWorld(t *testing.T) {
	value := World()
	if value != "Foo::World" {
		t.Fatal("Got unexpected value for Foo::World")
	}
}

//func TestBaz(t *testing.T) {
//value := Baz()
//if value != "Foo::Baz" {
//t.Fatal("Got unexpected value for Foo::Baz")
//}
//}
