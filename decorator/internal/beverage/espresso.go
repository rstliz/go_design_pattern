package beverage

type Espresso struct {
}

func (b *Espresso) Cost() float64 {
	return 1.99
}

func (b *Espresso) Description() string {
	return "エスプレッソ"
}
