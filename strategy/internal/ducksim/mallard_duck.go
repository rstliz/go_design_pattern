package ducksim

import "fmt"

type MallardDuck struct {
	DuckBehavior
}

func NewMallardDuck() *MallardDuck {
	d := &MallardDuck{
		DuckBehavior{
			quackBehavior: Quack{},
			flyBehavior:   FlyWithWings{},
		},
	}
	return d
}

func (d MallardDuck) Display() {
	fmt.Println("本物のマガモです")
}
