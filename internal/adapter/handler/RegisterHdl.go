package handler

import (
	"collection/errs"
	"collection/internal/core/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type registerHdl struct {
	svc domain.RegisterSvc
}

func NewRegisterHdl(svc domain.RegisterSvc) registerHdl {
	return registerHdl{
		svc: svc,
	}
}

func (h registerHdl) GetRegisters(c *gin.Context) {
	res, err := h.svc.GetAllRegister()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h registerHdl) GetRegister(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("registerID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := h.svc.GetRegister(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h registerHdl) AddRegister(c *gin.Context) {
	req := domain.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	res, err := h.svc.AddRegister(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h registerHdl) UpdateRegister(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("registerID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.RegisterReq{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateRegister(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update register success!!",
	})
}
func (h registerHdl) UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("registerID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.RoleId{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateRole(id, req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update register success!!",
	})
}
func (h registerHdl) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("registerID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	req := domain.RecordStatus{}
	err = c.BindJSON(&req)
	err = h.svc.UpdateStetus(id ,req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update register success!!",
	})
}


func (h registerHdl) DeleteRegister(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("registerID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err = h.svc.DeleteRegister(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete register success!!",
	})
}
func (h registerHdl) Login(ctx *gin.Context) {
	loginRequest := domain.LoginReq{}
	err := ctx.ShouldBindJSON(&loginRequest)
	errs.ErrorPanic(err)

	token, err_token := h.svc.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := domain.Resp{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	user, _ := h.svc.GetProfile(loginRequest.Username)
	resp := domain.LoginResp{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := domain.Resp{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Token:   resp.Token,
		User:    user,
	}
	// ctx.SetCookie("token", token, config.TokenMaxAge60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (h registerHdl) GetProfile(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(string)
	user, _ := h.svc.GetProfile(currentUser)
	webResponse := domain.Resp{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		User:    user,
	}
	c.JSON(http.StatusOK, webResponse)

}
