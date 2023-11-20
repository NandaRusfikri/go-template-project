package controller

import (
	"github.com/gin-gonic/gin"
	"go-template-project/dto"
	"go-template-project/module/user"
	"go-template-project/pkg"
	"go-template-project/util"
	"net/http"
)

type UserControllerHTTP struct {
	user_usecase user.UserUseCase
}

func InitUserControllerHTTP(route *gin.Engine, user_usecase user.UserUseCase) {

	controller := &UserControllerHTTP{
		user_usecase: user_usecase,
	}
	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/users", pkg.Auth, controller.UserList)
	groupRoute.POST("/user", pkg.Auth, controller.InsertUser)
	groupRoute.PUT("/user", pkg.Auth, controller.UserUpdate)
	groupRoute.POST("/user/change-password", pkg.Auth, controller.ChangePassword)
}

// ChangePassword
// @Tags User
// @Summary Change Password
// @Description API for change password user
// @ID User-ChangePassword
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body dto.ChangePassword true "body data"
// @Success 200
// @Router /v1/user/change-password [post]
func (c *UserControllerHTTP) ChangePassword(ctx *gin.Context) {

	var input dto.ChangePassword

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	err := c.user_usecase.ChangePasswordService(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "Change Password Success", http.StatusOK, 0, nil)
	}
}

// UserList
// @Tags User
// @Summary  User List
// @Description  API untuk mengambil data list user
// @ID User-UserList
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body dto.UsersRequest true "body data"
// @Success 200 {array}  dto.UsersResponse
// @Router /v1/users [get]
func (c *UserControllerHTTP) UserList(ctx *gin.Context) {

	var input dto.UsersRequest

	if err := ctx.ShouldBindQuery(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid", 400, 0, nil)
		return
	}
	res, count, err := c.user_usecase.ListUserService(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "Get Users Success", http.StatusOK, count, res)
	}
}

// InsertUser
// @Tags User
// @Summary  User Insert
// @Description  API untuk menambahkan user baru
// @ID User-UserInsert
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body dto.UserInsert true "body data"
// @Success 200
// @Router /v1/user [post]
func (c *UserControllerHTTP) InsertUser(ctx *gin.Context) {

	var input dto.UserInsert

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	err := c.user_usecase.CreateNewUserService(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "User Insert Success", http.StatusOK, 0, nil)
	}

}

// UserUpdate
// @Tags User
// @Summary  User Update
// @Description  API untuk mengedit data user
// @ID User-UserUpdate
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body dto.UserUpdate true "body data"
// @Success 200
// @Router /v1/user [put]
func (c *UserControllerHTTP) UserUpdate(ctx *gin.Context) {

	var input dto.UserUpdate

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "Request Invalid "+err.Error(), 400, 0, nil)
		return
	}
	res, err := c.user_usecase.UpdateUserService(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "User Update Success", http.StatusOK, 0, res)
	}
}
