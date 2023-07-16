package main

import "fmt"

// Struct embedding

type Position struct {
	posx int
	posy int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func main() {
	e := &SpecialEntity{
		specialField: 89.99,
		Entity: Entity{
			// name: "my special entity",
			// id:      "id 1 special",
			version: "version 1.1 special",
			Position: Position{
				posx: 100,
				posy: 200,
			},
		},
	}
	e.id = "id 1 special"
	e.name = "foo"
	e.posx = 69
	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", e.Entity.name)
	fmt.Printf("%+v\n", e.name)
}
