package util

import (
	"github.com/gin-gonic/gin"
	"go-template-project/schemas"
	_ "image/jpeg"
	_ "image/png"
	"reflect"

	"strings"
)

func SplitOrderQuery(order string) (string, string) {
	orderA := strings.Split(order, "|")
	if len(orderA) > 1 {
		if orderA[0] != "" && orderA[1] != "" {
			return orderA[0], orderA[1]
		}
	}

	return orderA[0], "ASC"
}

func MapStruct(source interface{}, dest interface{}) {
	valueOfSource := reflect.ValueOf(source)
	valueOfDest := reflect.ValueOf(dest)

	if valueOfSource.Kind() != reflect.Struct || valueOfDest.Kind() != reflect.Ptr {
		panic("Tipe data tidak valid")
	}

	typeOfSource := valueOfSource.Type()
	typeOfDest := valueOfDest.Elem().Type()

	for i := 0; i < typeOfSource.NumField(); i++ {
		fieldName := typeOfSource.Field(i).Name
		destField, found := typeOfDest.FieldByName(fieldName)
		if found {
			destValueField := valueOfDest.Elem().FieldByName(destField.Name)
			destValueField.Set(valueOfSource.FieldByName(fieldName))
		}
	}
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Count int64, Data interface{}) {

	jsonResponse := schemas.ResponseSuccess{
		Code: StatusCode,
		//Method:  Method,
		Count:   Count,
		Message: Message,
		Data:    Data,
		//Items:   Items,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}
