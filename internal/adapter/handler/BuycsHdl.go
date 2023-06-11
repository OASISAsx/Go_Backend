package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type buycsHdl struct {
	svc domain.BuycsSvc
}

func NewBuycsHdl(svc domain.BuycsSvc) buycsHdl {
	return buycsHdl{
		svc: svc,
	}
}

func (h buycsHdl) GetBuycss(c *gin.Context) {
	res, err := h.svc.GetAllBuycs()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h buycsHdl) GetBuycs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("buycsID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetBuycs(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h buycsHdl) AddBuycs(c *gin.Context) {
	req := domain.BuycsRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddBuycs(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h buycsHdl) UpdateBuycs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("buycsID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.BuycsRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateBuycs(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update buycs success!!",
	})
}

func (h buycsHdl) DeleteBuycs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("buycsID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteBuycs(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete buycs success!!",
	})
}
