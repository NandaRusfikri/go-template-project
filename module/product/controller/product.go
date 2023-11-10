package controller

import (
	"github.com/gin-gonic/gin"
	"go-template-project/module/product"
	"go-template-project/schemas"
	"go-template-project/util"
	"net/http"
)

type ProductControllerHTTP struct {
	product_usecase product.ItemUseCase
}

func InitProductControllerHTTP(route *gin.Engine, service product.ItemUseCase) {

	controller := &ProductControllerHTTP{
		product_usecase: service,
	}
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/products", controller.ProductList)
}

// ProductList
// @Tags Product
// @Summary Product List
// @Description Product List
// @ID Item-ProductList
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.ProductsRequest true "body data"
// @Success 200
// @Router /v1/products [post]
func (c *ProductControllerHTTP) ProductList(ctx *gin.Context) {

	var input schemas.ProductsRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		util.APIResponse(ctx, "request invalid", 400, 0, nil)
		return
	}
	res, count, err := c.product_usecase.ProductList(input)

	if err.Error != nil {
		util.APIResponse(ctx, err.Error.Error(), err.Code, 0, nil)
	} else {
		util.APIResponse(ctx, "List Success", http.StatusOK, count, res)
	}
}
