package strategy

import "fmt"

type MallardDuck struct {
	Duck
	flyBehavior   FlyWithWings
	quackBehavior Quack
}

func (d Duck) Display() {
	fmt.Println("本物のマガモです。")
}
