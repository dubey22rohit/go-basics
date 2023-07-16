package main

import (
	"fmt"
	"strings"
)

// Pure functions are good when no state is required
// When state is required use structs and / or interfaces

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("hello Sailor!", Uppercase))
	fmt.Println(transformString("hello Sailor!", Prefixer("FOO_")))
	fmt.Println(transformString("hello Sailor!", Prefixer("BAR_")))
}
