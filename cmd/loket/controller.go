package loket

import (
	"antrian/utils"
	cstmvalidation "antrian/utils/cstmValidation"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func Get(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)
	testInsert := LoketTable{
		NamaLoket:   "test",
		StatusLoket: "1",
		CreatedAt:   time.Now(),
	}

	db.Create(&testInsert)
	fmt.Println(db)
	// db := database.DBMySQLConnecting()
	// db.create
	ctx.JSON(200, gin.H{
		"test": "helo",
	})
}

func CreateNew(ctx *gin.Context) {

	var create Create
	validate := validator.New()
	validate.RegisterValidation("noWhiteSpace", cstmvalidation.NoWhiteSpace)
	validate.RegisterValidation("noDot", cstmvalidation.NoDot)

	if err := ctx.ShouldBindJSON(&create); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"type":    "Unmarshal Payload",
			"message": err.Error(),
		})
		return
	}

	if err := validate.Struct(create); err != nil {
		fieldErrors := utils.ValidateReflector(create, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"type":    "validation",
			"message": "error",
			"data":    fieldErrors,
		})
		return
	}

	db := utils.InitConnection().DB
	testInsert := LoketTable{
		NamaLoket:   "test",
		StatusLoket: "1",
		CreatedAt:   time.Now(),
	}

	db.Create(&testInsert)
	ctx.JSON(200, gin.H{
		"test": "helo",
	})
}
