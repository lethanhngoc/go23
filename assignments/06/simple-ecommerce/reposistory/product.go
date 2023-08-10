package reposistory

type Product struct {
	ProductId    int64  `json:"product_id" gorm:"column:product_id;"`
	ProductTitle string `json:"product_title" gorm:"column:product_title;""`
	ProductPrice int64  `json:"product_price" gorm:"column:product_price;"`
}

func NewProduct(productId int64, productTitle string, productPrice int64) *Product {
	return &Product{ProductId: productId, ProductTitle: productTitle, ProductPrice: productPrice}
}

func (item Product) GetProductId() int64 {
	return item.ProductId
}

func (item Product) GetProductTitle() string {
	return item.ProductTitle
}

func (item Product) GetProductPrice() int64 {
	return item.ProductPrice
}
