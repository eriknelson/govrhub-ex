package foo

import (
	"fmt"
)

func Hello() string {
	fmt.Println("Test")
	return "Foo::Hello"
}

func World() string {
	return "Foo::World"
}
