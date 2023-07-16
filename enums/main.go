package main

import "fmt"

type WeaponType int

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "Sword"
	case Axe:
		return "Axe"
	}

	return ""
}

const (
	Axe WeaponType = iota
	Sword
	WoodenStick
	Knife
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 10
	case Knife:
		return 40
	default:
		panic("Weapon does not exist")
	}
}

func main() {
	//Type WeaponType has a String() function attached to it, so it
	// is possible to %s and get string repr of it
	fmt.Printf("damage of weapon (%s): (%d)\n", Axe, getDamage(Axe))
	fmt.Printf("damage of weapon (%s): (%d)\n", Sword, getDamage(Sword))
	fmt.Printf("damage of weapon (%s): (%d)\n", WoodenStick, getDamage(WoodenStick))
}
