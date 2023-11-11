package schemas

type ProductsRequest struct {
	SearchText string `json:"search_text" form:"search_text" example:"Search name ku"`
	OrderField string `json:"order_field" form:"order_field" example:"id|desc"`
	Page       int    `json:"page"  example:"1"`
	Limit      int    `json:"limit" example:"10"`
}
