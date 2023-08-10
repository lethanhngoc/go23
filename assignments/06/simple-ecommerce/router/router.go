package router

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	//router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("")
	{
		productGroup := v1.Group("products")
		{
			productController := new(controllers.ProductController)
			productGroup.POST("", productController.CreateProduct)
			productGroup.PUT("/:id", productController.UpdateProduct)
			productGroup.DELETE("/:id", productController.DeleteProductById)
			//productGroup.GET("/:id", productController.GetPro)
			productGroup.GET("", productController.GetProducts)
			productGroup.GET("/:id", productController.GetProductById)
		}
	}
	return router
}
