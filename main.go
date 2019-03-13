package main

import (
	"fmt"
)

func foo(v string) string {
	return "*" + v + "*"
}

func main() {
	fmt.Println(foo("hello"))
}
