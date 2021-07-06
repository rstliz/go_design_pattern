package main

import (
	"github.com/rstliz/strategy"
)

func main() {
	performDuck(strategy.NewMallardDuck())
}

func performDuck(d strategy.Duck) {
	d.Display()
	d.Quack()
}
