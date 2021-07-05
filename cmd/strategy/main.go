package main

import (
	"github.com/rstliz/strategy"
)

func main() {
	mallard := &strategy.MallardDuck{}
	mallard.Display()
	mallard.PerformQuack()
	mallard.PerformFly()
}
