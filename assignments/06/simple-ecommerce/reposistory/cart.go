package reposistory

type Cart struct {
	CartId     int64     `json:"id" gorm:"column:cart_id;"`
	TotalPrice int64     `json:"total_price" gorm:"column:total_price;"`
	Products   []Product `json:"list_product" gorm:"column:list_product;"`
}
