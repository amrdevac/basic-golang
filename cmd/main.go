package main

import (
	"antrian/cmd/loket"
	"antrian/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InjectDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {
	utils.InitConnection()

	// db := database.DBMySQLConnecting()

	router := gin.Default()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// route.Use(InjectDB(db))
	router.GET("/loket", loket.Get)
	router.POST("/loket", loket.CreateNew)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  0 * time.Millisecond,
		WriteTimeout: 1 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	// route.Run() // listen and serve on 0.0.0.0:8080
}
