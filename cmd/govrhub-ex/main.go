package main

import (
	"fmt"
	"github.com/eriknelson/govrhub-ex/pkg/bar"
	"github.com/eriknelson/govrhub-ex/pkg/foo"
)

func main() {
	fmt.Println(foo.Hello())
	fmt.Println(foo.World())
	fmt.Println(bar.Hello())
}
