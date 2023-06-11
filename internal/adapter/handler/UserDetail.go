package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userdetailHdl struct {
	svc domain.UserDetailSvc
}

func NewUserDetailHdl(svc domain.UserDetailSvc) userdetailHdl {
	return userdetailHdl{
		svc: svc,
	}
}

func (h userdetailHdl) GetUserDetails(c *gin.Context) {
	res, err := h.svc.GetAllUserDetail()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h userdetailHdl) GetUserDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userdetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetUserDetail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h userdetailHdl) AddUserDetail(c *gin.Context) {
	req := domain.UserDetailRequest{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddUserDetail(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h userdetailHdl) UpdateUserDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userdetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.UserDetailRequest{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateUserDetail(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update userdetail success!!",
	})
}

func (h userdetailHdl) DeleteUserDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userdetailID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteUserDetail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete userdetail success!!",
	})
}
