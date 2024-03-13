package product

import "testing"

func TestGetDiscountPrice(t *testing.T) {
	datas := []struct {
		product  Product
		expected float64
	}{
		{
			Product{
				Name:     "Kopi",
				Price:    100,
				Discount: 0.1,
			},
			90, // Expected result
		},
		{
			Product{
				Name:     "Teh",
				Price:    50,
				Discount: 0,
			},
			50, // Expected Result
		},
	}

	for _, d := range datas {
		result := d.product.GetDiscountPrice()
		if result != d.expected {
			t.Errorf("For product %s with price %f expected %f but got %f",
				d.product.Name, d.product.Price, d.expected, result)
		}
	}
}
