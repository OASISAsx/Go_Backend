package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type areaapHdl struct {
	svc domain.AreaapSvc
}

func NewAreaapHdl(svc domain.AreaapSvc) areaapHdl {
	return areaapHdl{
		svc: svc,
	}
}

func (h areaapHdl) GetAreaaps(c *gin.Context) {
	res, err := h.svc.GetAllAreaap()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h areaapHdl) GetAreaap(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areaapID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetAreaap(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h areaapHdl) AddAreaap(c *gin.Context) {
	req := domain.AreaapRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddAreaap(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h areaapHdl) UpdateAreaap(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areaapID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.AreaapRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateAreaap(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update areaap success!!",
	})
}

func (h areaapHdl) DeleteAreaap(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areaapID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteAreaap(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete areaap success!!",
	})
}
