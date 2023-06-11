package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type areapvHdl struct {
	svc domain.AreapvSvc
}

func NewAreapvHdl(svc domain.AreapvSvc) areapvHdl {
	return areapvHdl{
		svc: svc,
	}
}

func (h areapvHdl) GetAreapvs(c *gin.Context) {
	res, err := h.svc.GetAllAreapv()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h areapvHdl) GetAreapv(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areapvID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetAreapv(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h areapvHdl) AddAreapv(c *gin.Context) {
	req := domain.AreapvRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddAreapv(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h areapvHdl) UpdateAreapv(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areapvID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.AreapvRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateAreapv(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update areapv success!!",
	})
}

func (h areapvHdl) DeleteAreapv(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areapvID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteAreapv(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete areapv success!!",
	})
}
