package main

import "fmt"

//A pointer is an 8 byte long integer

type Player struct {
	HP int
}

// if player is not a pointer we are adjusting the copy of the player
// not the actual player
func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. New Damage-> ", player.HP)
}

func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage. New Damage-> ", p.HP)
}

type BigData struct {
	//1 gb of memory
}

func doSomethingWithData(data *BigData) {
	//manipulate the data inside this function
}

func main() {
	player := &Player{
		HP: 100,
	}

	TakeDamage(player, 10)
	player.TakeDamage(40)
	fmt.Printf("%+v\n", player)

	data := &BigData{} //1gb

	//copy 1gb of data inside of that function
	//Not efficient, we can use pointers
	// doSomethingWithData(data)
	doSomethingWithData(data) //Now passing pointer, copying only 8 bytes of memory
}
