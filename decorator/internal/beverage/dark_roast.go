package beverage

type DarkRoast struct {
}

func (b *DarkRoast) Cost() float64 {
	return .99
}

func (b *DarkRoast) Description() string {
	return "ダークロースト"
}
