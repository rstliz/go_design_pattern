package strategy

import "fmt"

type FlyWithWings struct {
}

func (f FlyWithWings) Fly() {
	fmt.Println("飛んでいます！")
}
