package controller

import (
	"github.com/gin-gonic/gin"
	"go-template-project/dto"
	"go-template-project/module/product"
	"go-template-project/util"
	"net/http"
)

type ProductControllerHTTP struct {
	productUsecase product.UseCaseInterface
}

func InitProductControllerHTTP(route *gin.Engine, service product.UseCaseInterface) {

	controller := &ProductControllerHTTP{
		productUsecase: service,
	}
	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/products", controller.ProductList)
}

// ProductList
//
//	@Tags			Product
//	@Summary		Product List
//	@Description	Product List
//	@ID				Item-GetList
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			request query	dto.ProductsRequest true "as"
//	@Success		200
//	@Router			/v1/products [get]
func (c *ProductControllerHTTP) ProductList(ctx *gin.Context) {

	var input dto.ProductsRequest

	if err := ctx.ShouldBindQuery(&input); err != nil {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "request invalid " + err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	res, count, err := c.productUsecase.GetList(input)

	if err.Error != nil {
		util.APIResponse(ctx, dto.APIResponse{Message: err.Error.Error(), Code: err.Code})
	} else {
		util.APIResponse(ctx, dto.APIResponse{
			Message: "List Success",
			Code:    http.StatusOK,
			Count:   count,
			Data:    res,
		})
	}
}
