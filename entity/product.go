package entity

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

func (p Product) ProdukStatus() string {
	var status string
	if p.Stock < 3 {
		status = "produk hampir habis"
	} else if p.Stock < 10 {
		status = "produk terbatas"
	}
	return status
}
