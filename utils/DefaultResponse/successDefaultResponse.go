package defaultresponse

import (
	"antrian/utils"

	"github.com/gin-gonic/gin"
)

func GetListResponseSuccess(ctx *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "Success Load"
	}
	ctx.JSON(200, utils.BaseTypeResponse{
		Status:  true,
		Type:    "success",
		Message: message,
		Data:    data,
	})

}
func InsertResponseSuccess(ctx *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "Success Insert data"
	}
	ctx.JSON(200, utils.BaseTypeResponse{
		Status:  true,
		Type:    "success",
		Message: message,
		Data:    data,
	})
}

func UpdateResponseSuccess(ctx *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "Success update data"
	}
	ctx.JSON(200, utils.BaseTypeResponse{
		Status:  true,
		Type:    "success",
		Message: message,
		Data:    data,
	})
}
