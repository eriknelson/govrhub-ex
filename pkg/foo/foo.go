package foo

import (
	"fmt"
)

func Hello() string {
	fmt.Println("Test")
	return "Foo::Hello"
}

func World() string {
	fmt.Println("Hello")
	return "Foo:World"
}
