package dto

import (
	"errors"
	"mime/multipart"
)

type ResponseBaseDto struct {
	Code         int    `json:"code"`
	Message      string `json:"message,omitempty"`
	MessageError string `json:"message_error,omitempty"`
}

type CallAPI struct {
	Method       string
	Url          string
	ContentType  string
	Headers      map[string]interface{}
	BodyRequest  string
	BodyResponse string
	HttpCode     int
}

func (d *CallAPI) Validate() error {
	if d.Method == "" {
		return errors.New("method required")
	}
	if d.Url == "" {
		return errors.New("url required")
	}

	return nil
}

type ExportBase64ResponseDto struct {
	ResponseBaseDto
	Extension string `json:"extension"`
	Base64    string `json:"base_64"`
}
type UploadFileRequestDto struct {
	File []*multipart.FileHeader `json:"-" form:"file" binding:"required"`
	Path string                  `json:"path"`
}
type UploadFileResponseDto struct {
	ResponseBaseDto
	Count int64                   `json:"count,omitempty"`
	Items []*UploadFileRequestDto `json:"items,omitempty"`
}
