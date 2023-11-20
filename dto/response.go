package dto

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Count   int64       `json:"count,omitempty"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code  int
	Error error
}
