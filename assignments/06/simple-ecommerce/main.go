package main

import (
	"shopping-cart/router"
)

func main() {

	r := router.NewRouter()
	r.Run(":8080")

	//router := mux.NewRouter()
	//router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
	//	productId, _ := strconv.ParseInt(r.URL.Query().Get("product_id"), 10, 64)
	//	productTitle := r.URL.Query().Get("product_title")
	//	productPrice, _ := strconv.ParseInt(r.URL.Query().Get("product_price"), 10, 64)
	//	product := model.NewProduct(productId, productTitle, productPrice)
	//	if err := db.Create(&product).Error; err != nil {
	//		log.Fatalln("error: ", err)
	//	}
	//}).Methods("POST")
	//
	//router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
	//	productId, _ := strconv.ParseInt(r.URL.Query().Get("product_id"), 10, 64)
	//	productTitle := r.URL.Query().Get("product_title")
	//	productPrice, _ := strconv.ParseInt(r.URL.Query().Get("product_price"), 10, 64)
	//	product := model.NewProduct(productId, productTitle, productPrice)
	//	if err := db.Create(&product).Error; err != nil {
	//		log.Fatalln("error: ", err)
	//	}
	//}).Methods("POST")
	//log.Fatal(http.ListenAndServe(":8080", router))
}

func createItem() {

}
