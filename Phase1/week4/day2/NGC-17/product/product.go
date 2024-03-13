package product

type Product struct {
	ID    int
	Name  string
	Price float64
}

var p []Product

func (product *Product) AddProduct() []Product {
	p = append(p, Product{
		Name:  product.Name,
		Price: product.Price,
	})

	return p
}

var p2 = []Product{
	{ID: 1, Name: "Kopi", Price: 100},
}

func GetProductById(id int) Product {
	for i, v := range p2 {
		if v.ID == id {
			return p2[i]
		}
	}
	return Product{}
}

func (product *Product) UpdateProduct() Product {
	p2[0].ID = product.ID
	p2[0].Name = product.Name
	p2[0].Price = product.Price
	return p2[0]
}

func DeleteProductById(id int) Product {
	for i, v := range p2 {
		if v.ID == id {
			p2[i].ID = -1
			p2[i].Name = ""
			p2[i].Price = -1
			return p2[i]
		}
	}
	return Product{}
}
