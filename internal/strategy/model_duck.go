package strategy

import "fmt"

type ModelDuck struct {
	Duck
}

func (d ModelDuck) Display() {
	fmt.Println("模型の鴨です")
}
