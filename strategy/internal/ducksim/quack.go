package ducksim

import "fmt"

type Quack struct {
	QuackBehavior
}

func (q Quack) Quack() {
	fmt.Println("ガーガー！")
}
