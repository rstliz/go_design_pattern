package ducksim

import "fmt"

type FlyWithWings struct {
	FlyBehavior
}

func (f FlyWithWings) Fly() {
	fmt.Println("飛んでいます！")
}
