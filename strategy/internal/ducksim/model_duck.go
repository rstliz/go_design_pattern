package ducksim

import "fmt"

type ModelDuck struct {
	DuckBehavior
}

func NewModelDuck() *ModelDuck {
	d := &ModelDuck{
		DuckBehavior{
			quackBehavior: MuteQuack{},
			flyBehavior:   FlyNoWay{},
		},
	}
	return d
}

func (d ModelDuck) Display() {
	fmt.Println("模型の鴨です")
}
