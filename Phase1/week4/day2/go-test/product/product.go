package product

type Product struct {
	Name     string
	Price    float64
	Discount float64
}

func (p *Product) GetDiscountPrice() float64 {
	return p.Price * (1 - p.Discount)
}