package dto

type ProductsRequest struct {
	SearchText string `form:"search_text" example:"Search name ku"`
	OrderField string `form:"order_field" example:"id|desc"`
	Page       int    `form:"page"  example:"1"`
	Limit      int    `form:"limit" example:"10"`
}
