package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type customerHdl struct {
	svc domain.CustomerSvc
}

func NewCustomerHdl(svc domain.CustomerSvc) customerHdl {
	return customerHdl{
		svc: svc,
	}
}

func (h customerHdl) GetCustomers(c *gin.Context) {
	res, err := h.svc.GetAllCustomer()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h customerHdl) GetCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetCustomer(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h customerHdl) AddCustomer(c *gin.Context) {
	req := domain.CustomerReq{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddCustomer(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h customerHdl) UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.CustomerReq{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateCustomer(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update customer success!!",
	})
}

func (h customerHdl) DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteCustomer(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete customer success!!",
	})
}
