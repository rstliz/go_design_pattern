package strategy

type Duck interface {
	Display()
	QuackBehavior
}

type DuckBehavior struct {
	quackBehavior QuackBehavior
}

func (d DuckBehavior) Quack() {
	d.quackBehavior.Quack()
}

//func (d Duck) Swim() {
//	fmt.Println("すべての鴨は浮かびます。おとりの鴨でも！")
//}
