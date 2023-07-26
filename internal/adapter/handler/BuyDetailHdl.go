package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type buydetailHdl struct {
	svc domain.BuyDetailSvc
}

func NewBuyDetailHdl(svc domain.BuyDetailSvc) buydetailHdl {
	return buydetailHdl{
		svc: svc,
	}
}

func (h buydetailHdl) GetBuyDetails(c *gin.Context) {
	res, err := h.svc.GetAllBuyDetail()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h buydetailHdl) SearchBuyDetail(c *gin.Context) {
	name := c.Param("name")

	res, _ := h.svc.SearchBuyDetail(name)
	c.JSON(http.StatusOK, res)
}

func (h buydetailHdl) GetBuyDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("buydetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetBuyDetail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h buydetailHdl) AddBuyDetail(c *gin.Context) {
	req := domain.BuyDetailRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddBuyDetail(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h buydetailHdl) UpdateBuyDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("buydetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.BuyDetailRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateBuyDetail(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update buydetail success!!",
	})
}

func (h buydetailHdl) DeleteBuyDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("buydetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteBuyDetail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete buydetail success!!",
	})
}
func (h buydetailHdl) GetAllBuyDetailId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("buydetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetAllBuyDetailId(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h buydetailHdl) GetAllUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("UserId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetAllUserId(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}