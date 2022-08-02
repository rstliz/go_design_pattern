package beverage

type Milk struct {
	beverage Beverage
}

func NewMilk(beverage Beverage) *Milk {
	b := &Milk{
		beverage: beverage,
	}
	return b
}

func (b *Milk) Cost() float64 {
	return .10 + b.beverage.Cost()
}

func (b *Milk) Description() string {
	return b.beverage.Description() + ". ミルク"
}
