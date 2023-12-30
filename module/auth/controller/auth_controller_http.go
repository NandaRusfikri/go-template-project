package controller

import (
	"github.com/gin-gonic/gin"
	"go-template-project/dto"
	"go-template-project/module/auth"
	"go-template-project/util"
	"net/http"
)

type AuthControllerHTTP struct {
	authUsecase auth.UseCaseInterface
}

func InitAuthControllerHTTP(route *gin.Engine, authUsecase auth.UseCaseInterface) {

	controller := &AuthControllerHTTP{
		authUsecase: authUsecase,
	}
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/auth/login", controller.Login)
	groupRoute.POST("/auth/forgot-password", controller.ForgotPassword)
	groupRoute.POST("/auth/reset-password", controller.ResetPassword)
}

// ResetPassword
//
//	@Tags			Auth
//	@Summary		Reset Password
//	@Description	API for confirm reset password
//	@ID				User-ResetPassword
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			data	body	dto.ResetPassword	true	"body data"
//	@Success		200
//	@Router			/v1/auth/reset-password [post]
func (c *AuthControllerHTTP) ResetPassword(ctx *gin.Context) {

	var input dto.ResetPassword

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	err := c.authUsecase.ResetPassword(input)

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{Message: err.Error.Error(), Code: err.Code})
	} else {
		util.APIResponse(ctx, dto.APIResponse{Message: "Reset Password Success", Code: http.StatusOK})
	}
}

// ForgotPassword
//
//	@Tags			Auth
//	@Summary		Forgot Password
//	@Description	API for Request Forgot Password
//	@ID				User-ForgotPassword
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			data	body	dto.ForgotPassword	true	"body data"
//	@Success		200
//	@Router			/v1/auth/forgot-password [post]
func (c *AuthControllerHTTP) ForgotPassword(ctx *gin.Context) {

	var input dto.ForgotPassword

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	err := c.authUsecase.ForgotPassword(input)

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{Message: err.Error.Error(), Code: err.Code})
	} else {
		util.APIResponse(ctx, dto.APIResponse{Message: "Request Forgot Password Success", Code: http.StatusOK})
	}
}

// Login
//
//	@Tags			Auth
//	@Summary		Login
//	@Description	API for Login
//	@ID				User-Login
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			data	body	dto.LoginRequest	true	"body data"
//	@Success		200
//	@Router			/v1/auth/login [post]
func (c *AuthControllerHTTP) Login(ctx *gin.Context) {

	var input dto.LoginRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	res, err := c.authUsecase.Login(input)

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{Message: err.Error.Error(), Code: err.Code})
	} else {
		util.APIResponse(ctx, dto.APIResponse{Message: "Login Success", Code: http.StatusOK, Data: res})
	}
}
