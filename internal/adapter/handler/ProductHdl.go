package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHdl struct {
	svc domain.ProductSvc
}

func NewProductHdl(svc domain.ProductSvc) productHdl {
	return productHdl{
		svc: svc,
	}
}

func (h productHdl) GetProducts(c *gin.Context) {
	res, err := h.svc.GetAllProduct()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h productHdl) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetById(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h productHdl) AddProduct(c *gin.Context) {
	req := domain.ProductRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddProduct(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h productHdl) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.ProductRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateProduct(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update product success!!",
	})
}

func (h productHdl) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteProduct(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete product success!!",
	})
}
func (h productHdl) Search(c *gin.Context) {
	productname := c.Param("ProductName")

	res, _ := h.svc.Search(productname)
	c.JSON(http.StatusOK, res)
}

func (h *productHdl) UpdateStatusProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	req := domain.StatusProduct{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateStatusProduct(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product success!!",
	})

}
func (h *productHdl) UpdateSellStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ProductId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	req := domain.SellStatusProduct{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateSellStatus(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product success!!",
	})

}
func (h *productHdl) UpdateCount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ProductId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	req := domain.ProductRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateCount(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Product success!!",
	})

}