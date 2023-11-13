package controller

import (
	"github.com/gin-gonic/gin"
	"go-template-project/module/auth"
	"go-template-project/schemas"
	"go-template-project/util"
	"net/http"
)

type AuthControllerHTTP struct {
	auth_usecase auth.AuthUseCase
}

func InitAuthControllerHTTP(route *gin.Engine, auth_usercase auth.AuthUseCase) {

	controller := &AuthControllerHTTP{
		auth_usecase: auth_usercase,
	}
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/auth/login", controller.LoginController)
	groupRoute.POST("/auth/forgot-password", controller.RequestForgotPassword)
	groupRoute.POST("/auth/reset-password", controller.ResetPassword)
}

// ResetPassword
// @Tags Auth
// @Summary Reset Password
// @Description API for confirm reset password
// @ID User-ResetPassword
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.ResetPassword true "body data"
// @Success 200
// @Router /v1/auth/reset-password [post]
func (c *AuthControllerHTTP) ResetPassword(ctx *gin.Context) {

	var input schemas.ResetPassword

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	err := c.auth_usecase.ResetPassword(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "Reset Password Success", http.StatusOK, 0, nil)
	}
}

// RequestForgotPassword
// @Tags Auth
// @Summary  Forgot Password
// @Description  API for Request Forgot Password
// @ID User-ForgotPassword
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.ForgotPassword true "body data"
// @Success 200
// @Router /v1/auth/forgot-password [post]
func (c *AuthControllerHTTP) RequestForgotPassword(ctx *gin.Context) {

	var input schemas.ForgotPassword

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	err := c.auth_usecase.RequestForgotPassword(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "Request Forgot Password Success", http.StatusOK, 0, nil)
	}
}

// LoginController
// @Tags Auth
// @Summary Login
// @Description API for Login
// @ID User-LoginController
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.LoginRequest true "body data"
// @Success 200
// @Router /v1/auth/login [post]
func (c *AuthControllerHTTP) LoginController(ctx *gin.Context) {

	var input schemas.LoginRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	res, err := c.auth_usecase.Login(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "Login Success", http.StatusOK, 0, res)
	}
}
