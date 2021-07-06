package ducksim

import "fmt"

type Duck interface {
	Display()
	Swim()
	QuackBehavior
	FlyBehavior
}

type DuckBehavior struct {
	quackBehavior QuackBehavior
	flyBehavior   FlyBehavior
}

func (d DuckBehavior) Quack() {
	d.quackBehavior.Quack()
}

func (d DuckBehavior) Fly() {
	d.flyBehavior.Fly()
}

func (d DuckBehavior) Swim() {
	fmt.Println("すべての鴨は浮かびます。おとりの鴨でも！")
}
