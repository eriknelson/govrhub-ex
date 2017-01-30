package foo

import (
	"fmt"
)

func Hello() string {
	return "Foo::Hello"
}

func World() string {
	fmt.Println("Hello")
	return "Foo:World"
}
