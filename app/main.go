package main

import (
	"fmt"

	tools "github.com/aj/toolkit"
)

func main() {
	var tools tools.Tools
	s := tools.RandomString(10)
	fmt.Println(s)
}
