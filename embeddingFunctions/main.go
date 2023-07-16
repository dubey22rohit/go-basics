package main

import "fmt"

type Position struct {
	x, y int
}

func (p *Position) Move(val int) {
	fmt.Println("the player moved by: ", val)
}

type Player struct {
	Position
}

func main() {
	p := Player{}
	//Even though Move is defined on Position, since Player embeds Position
	//Player can also use Move / inherit Move
	p.Move(100)
}
