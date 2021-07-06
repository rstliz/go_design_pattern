package ducksim

import "fmt"

type FlyNoWay struct {
	FlyBehavior
}

func (f FlyNoWay) Fly() {
	fmt.Println("飛べません")
}
