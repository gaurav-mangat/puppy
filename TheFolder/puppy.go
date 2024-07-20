package puppy

import (
	"fmt"

	"github.com/gaurav-mangat/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}
func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}
func From11(){
	fmt.Println("From version v4.0.0")
	
}
