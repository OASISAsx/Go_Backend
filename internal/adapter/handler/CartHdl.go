package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cartHdl struct {
	svc domain.CartSvc
}

func NewCartHdl(svc domain.CartSvc) cartHdl {
	return cartHdl{
		svc: svc,
	}
}

func (h cartHdl) GetCarts(c *gin.Context) {
	res, err := h.svc.GetAllCart()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h cartHdl) GetCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("CartID"))
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

func (h cartHdl) AddCart(c *gin.Context) {
	req := domain.CartRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddCart(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h cartHdl) UpdateCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("CartID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.CartRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateCart(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Cart success!!",
	})
}

func (h cartHdl) DeleteCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("CartID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteCart(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Cart success!!",
	})
}
