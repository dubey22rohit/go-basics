package main

import (
	"fmt"
	"go-anthony/modules/types"
	"go-anthony/modules/utils"
)

// User -> public access everywhere
// user -> private access BUT public in its own package

func main() {
	user := types.User{
		Username: utils.GetUsername(),
		Age:      utils.GetAge(),
	}
	// %v is anything
	// %+v : we can get the field names in the struct
	fmt.Printf("the user is %+v", user)
}
