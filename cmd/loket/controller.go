package loket

import (
	"antrian/utils"
	defaultresponse "antrian/utils/DefaultResponse"
	cstmvalidation "antrian/utils/cstmValidation"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Get(ctx *gin.Context) {
	offsetInt, errOffset := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if errOffset != nil {
		defaultresponse.NaNResponseError(ctx, "Offset")
		return
	}

	limitInt, errLimit := strconv.Atoi(ctx.DefaultQuery("limit", "2"))
	if errLimit != nil {
		defaultresponse.NaNResponseError(ctx, "Limit")
		return
	}

	getListLoket := GetListLoket(FilterGetList{
		NamaLoket: ctx.Query("nama_loket"),
		Status:    ctx.DefaultQuery("status_loket", "1"),
		Offset:    offsetInt,
		Limit:     limitInt,
	})
	if !getListLoket.Status {
		defaultresponse.QueryResponseError(ctx, getListLoket.Message)
		return
	}

	defaultresponse.GetListResponseSuccess(ctx, getListLoket.Data, "")
}

func CreateNew(ctx *gin.Context) {
	var create Create
	err := ctx.ShouldBindJSON(&create)
	fmt.Println(err)
	if err != nil {
		defaultresponse.BindingResponseError(ctx, err)
		return
	}

	validate := validator.New()
	validate.RegisterValidation("noWhiteSpace", cstmvalidation.NoWhiteSpace)
	validate.RegisterValidation("noDot", cstmvalidation.NoDot)
	if err := validate.Struct(create); err != nil {
		fieldErrors := utils.ValidateReflector(create, err)
		defaultresponse.ValidationResponseError(ctx, fieldErrors)
		return
	}

	create.InsertNewLoket()
	defaultresponse.InsertResponseSuccess(ctx, nil, "")
}

func UpdateStatus(ctx *gin.Context) {
	var updateDataStatus UpdateDataStatus
	err := ctx.ShouldBindJSON(&updateDataStatus)
	fmt.Println(err)
	if err != nil {
		defaultresponse.BindingResponseError(ctx, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(updateDataStatus); err != nil {
		fieldErrors := utils.ValidateReflector(updateDataStatus, err)
		defaultresponse.ValidationResponseError(ctx, fieldErrors)
		return
	}

	result := updateDataStatus.UpdateColStatus()
	if !result {
		defaultresponse.DataNotFoundResponseError(ctx)
		return
	}
	defaultresponse.UpdateResponseSuccess(ctx, updateDataStatus, "")
}
