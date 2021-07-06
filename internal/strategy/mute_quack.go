package strategy

import "fmt"

type MuteQuack struct {
	QuackBehavior
}

func (q MuteQuack) Quack() {
	fmt.Println("<沈黙>")
}
