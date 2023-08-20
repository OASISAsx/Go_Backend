package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type sellerDetailHdl struct {
	svc domain.SellerDetailSvc
}

func NewSellerDetailHdl(svc domain.SellerDetailSvc) sellerDetailHdl {
	return sellerDetailHdl{
		svc: svc,
	}
}

func (h sellerDetailHdl) GetSellerDetails(c *gin.Context) {
	res, err := h.svc.GetAllSellerDetail()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h sellerDetailHdl) GetSellerDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sellerdetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetSellerDetail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h sellerDetailHdl) AddSellerDetail(c *gin.Context) {
	req := domain.SellerDetailRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	res, err := h.svc.AddSellerDetail(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h sellerDetailHdl) UpdateSellerDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sellerdetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.SellerDetailRequest{}
	err = c.BindJSON(&req)
	if err != nil {
		c.Error(err)
	}
	err = h.svc.UpdateSellerDetail(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update SellerDetail success!!",
	})
}

func (h sellerDetailHdl) DeleteSellerDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sellerdetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteSellerDetail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete SellerDetail success!!",
	})
}
