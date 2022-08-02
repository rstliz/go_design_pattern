package beverage

type Whip struct {
	beverage Beverage
}

func NewWhip(beverage Beverage) *Whip {
	b := &Whip{
		beverage: beverage,
	}
	return b
}

func (b *Whip) Cost() float64 {
	return .10 + b.beverage.Cost()
}

func (b *Whip) Description() string {
	return b.beverage.Description() + ". ホイップ"
}
