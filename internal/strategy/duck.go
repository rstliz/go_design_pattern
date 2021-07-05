package strategy

import "fmt"

type Duck struct {
	Name          string
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d Duck) PerformFly() {
	d.flyBehavior.Fly()
}

func (d Duck) PerformQuack() {
	d.quackBehavior.Quack()
}

func (d Duck) Swim() {
	fmt.Println("すべての鴨は浮かびます。おとりの鴨でも！")
}
