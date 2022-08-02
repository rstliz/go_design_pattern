package beverage

type Mocha struct {
	beverage Beverage
}

func NewMocha(beverage Beverage) *Mocha {
	d := &Mocha{
		beverage: beverage,
	}
	return d
}

func (b *Mocha) Cost() float64 {
	return .20 + b.beverage.Cost()
}

func (b *Mocha) Description() string {
	return b.beverage.Description() + ". モカ"
}
