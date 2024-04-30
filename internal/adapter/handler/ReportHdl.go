package handler

import (
	"collection/internal/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type reportHdl struct {
	svc domain.ReportSvc
}

func NewReportHdl(svc domain.ReportSvc) reportHdl {
	return reportHdl{
		svc: svc,
	}
}

func (h reportHdl) GetReports(c *gin.Context) {
	res, err := h.svc.GetAllReport()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h reportHdl) GetReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("reportID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetReport(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h reportHdl) AddReport(c *gin.Context) {
	req := domain.ReportReq{}
	err := c.BindJSON(&req)
	res, err := h.svc.AddReport(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h reportHdl) UpdateReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("reportID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.ReportReq{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateReport(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update report success!!",
	})
}

func (h reportHdl) DeleteReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("reportID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteReport(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete report success!!",
	})
}
