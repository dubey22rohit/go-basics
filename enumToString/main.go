package main

import "fmt"

type Color int

// fmt.Stringer
// Color now implements the above interface, and whenever go tries to print the enum value
// It first checks for a String method, kinda like Python's __str__
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "Blue"
	case ColorBlack:
		return "Black"
	case ColorYellow:
		return "Yellow"
	case ColorPink:
		return "Pink"
	default:
		panic("invalid color input")
	}
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	fmt.Println("The color is: ", ColorBlue)
}
