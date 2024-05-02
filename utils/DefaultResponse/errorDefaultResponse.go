package defaultresponse

import (
	"antrian/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindingResponseError(ctx *gin.Context, err error) {
	ctx.JSON(400, utils.BaseTypeResponse{
		Status:  false,
		Type:    "BindingPayloadError",
		Message: err.Error(),
		Data:    nil,
	})
}

func DataNotFoundResponseError(ctx *gin.Context) {
	ctx.JSON(400, utils.BaseTypeResponse{
		Status:  false,
		Type:    "DataNotFound",
		Message: "Data Not Found",
		Data:    nil,
	})
}

func NaNResponseError(ctx *gin.Context, name string) {
	ctx.JSON(400, utils.BaseTypeResponse{
		Status:  false,
		Type:    "NaN",
		Message: name + " value not a number",
		Data:    nil,
	})
}

func QueryResponseError(ctx *gin.Context, QueryError string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  false,
		"type":    "QuerySQL",
		"message": QueryError,
		"data":    "",
	})
}

func ValidationResponseError(ctx *gin.Context, fieldErrors map[string]string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  false,
		"type":    "validation",
		"message": "error",
		"data":    fieldErrors,
	})
}
