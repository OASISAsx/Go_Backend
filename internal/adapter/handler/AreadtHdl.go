package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type areadtHdl struct {
	svc domain.AreadtSvc
}

func NewAreadtHdl(svc domain.AreadtSvc) areadtHdl {
	return areadtHdl{
		svc: svc,
	}
}

func (h areadtHdl) GetAreadts(c *gin.Context) {
	res, err := h.svc.GetAllAreadt()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h areadtHdl) GetAreadt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areadtID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetAreadt(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h areadtHdl) AddAreadt(c *gin.Context) {
	req := domain.AreadtRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddAreadt(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h areadtHdl) UpdateAreadt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areadtID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.AreadtRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateAreadt(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update areadt success!!",
	})
}

func (h areadtHdl) DeleteAreadt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("areadtID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteAreadt(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete areadt success!!",
	})
}
