package product

import (
	"testing"
)

func TestAddProduct(t *testing.T) {
	datas := []struct {
		product  []Product
		expected []Product
	}{
		{
			[]Product{
				{Name: "Kopi", Price: 100},
				{Name: "Susu", Price: 50},
			},
			[]Product{
				{Name: "Kopi", Price: 100},
				{Name: "", Price: 50},
			},
		},
	}

	for _, d := range datas {
		for i, v := range d.product {
			result := v.AddProduct()
			if result[i].Name != d.expected[i].Name || result[i].Price != d.expected[i].Price {
				t.Errorf("Result is {%s, %f} but the expected is {%s, %f}", result[i].Name, result[i].Price, d.expected[i].Name, d.expected[i].Price)
			}
		}
	}
}

func TestGetProductById(t *testing.T) {
	datas := []struct {
		id       int
		expected Product
	}{
		{
			1,
			Product{
				ID: 1, Name: "Kopi", Price: 100,
			},
		},
		{
			2,
			Product{
				ID: 1, Name: "Kopi", Price: 100,
			},
		},
	}

	for _, d := range datas {
		p := GetProductById(d.id)
		if p != d.expected {
			t.Errorf("Result ID: %d NOT FOUND", d.id)
		}
	}
}

func TestUpdateProduct(t *testing.T) {
	datas := []struct {
		updatedProduct Product
		expected       Product
	}{
		{
			Product{
				ID: 1, Name: "Susu", Price: 50,
			},
			Product{
				ID: 1, Name: "Susu", Price: 50,
			},
		},
		{
			Product{
				ID: 1, Name: "", Price: 100,
			},
			Product{
				ID: 1, Name: "Kopi", Price: 100,
			},
		},
	}

	for _, d := range datas {
		result := d.updatedProduct.UpdateProduct()
		if result != d.expected {
			t.Errorf("Update result is %v but the expected is %v", result, d.expected)
		}
	}
}

func TestDeleteProductById(t *testing.T) {
	datas := []struct {
		id int
	}{
		{
			1,
		},
		{
			2,
		},
	}

	var expected = Product{
		ID: -1, Name: "", Price: -1,
	}

	for _, d := range datas {
		p := DeleteProductById(d.id)
		if p != expected {
			t.Errorf("Result is: %v but the expected is %v", p, expected)
		}
	}
}
