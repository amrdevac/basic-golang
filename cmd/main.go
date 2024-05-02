package main

import (
	"antrian/cmd/loket"
	"antrian/utils"
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
	// route.Use(InjectDB(db))
	router.GET("/loket", loket.Get)
	router.POST("/loket", loket.CreateNew)
	router.PUT("/loket/status", loket.UpdateStatus)
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
