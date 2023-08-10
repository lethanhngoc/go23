package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/db"
	"shopping-cart/reposistory"
	"strconv"
)

func (p ProductController) CreateCart(c *gin.Context) {
	db := db.GetDB()
	var productItem reposistory.Product
	if err := c.ShouldBind(&productItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&productItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": productItem.ProductId})
}

func (p ProductController) GetAllCarts(c *gin.Context) {
	db := db.GetDB()
	type DataPaging struct {
		Page  int   `json:"page" form:"page"`
		Limit int   `json:"limit" form:"limit"`
		Total int64 `json:"total" form:"-"`
	}

	var paging DataPaging
	if err := c.ShouldBind(&paging); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if paging.Page <= 0 {
		paging.Page = 1
	}
	if paging.Limit <= 0 {
		paging.Limit = 10
	}
	offset := (paging.Page - 1) * paging.Limit
	var result []reposistory.Product
	if err := db.Table("carts").
		Count(&paging.Total).
		Offset(offset).
		Order("cart_id desc").
		Find(&result).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (p ProductController) GetCartById(c *gin.Context) {
	db := db.GetDB()
	var dataItem reposistory.Product
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("cart_id = ?", id).First(&dataItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dataItem})
}
