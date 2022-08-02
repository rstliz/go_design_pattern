package beverage

type HouseBlend struct {
}

func (b *HouseBlend) Cost() float64 {
	return .89
}

func (b *HouseBlend) Description() string {
	return "ハウスブレンド"
}
