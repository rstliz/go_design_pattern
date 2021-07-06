package main

import (
	"fmt"

	"github.com/rstliz/ducksim"
)

func main() {
	performDuck(ducksim.NewMallardDuck())
	performDuck(ducksim.NewModelDuck())
}

func performDuck(d ducksim.Duck) {
	d.Display()
	d.Quack()
	d.Fly()
	d.Swim()
	fmt.Println("")
}
