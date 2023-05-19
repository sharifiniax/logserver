package main

import (
//	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"loggermodule/core"
)

func main() {
	// ایجاد یک نمونه از ماژول لاگر


	// تنظیمات مربوط به فایل لاگ
	err := logger.SetLogFile("data/app.log", 10, 5, 30)
	if err != nil {
		// هندل خطا در صورتی که تنظیمات فایل لاگ انجام نشود
		panic(err)
	}

	defer logger.CloseLogFile()

	// تنظیمات مربوط به سطح لاگ
	logger.SetLogLevel(logger.Debug)

	// ایجاد روتر Gin
	router := gin.Default()

	// ایجاد یک روت برای لاگ گذاری
	router.POST("/logs", func(c *gin.Context) {
		var logData struct {
			Level   logger.LogLevel `json:"level"`
			Message string         `json:"message"`
			Tag     string         `json:"tag"`
		}

		// خواندن داده‌های لاگ از درخواست
		if err := c.ShouldBindJSON(&logData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// ارسال لاگ به ماژول لاگر
		logger.Log(logData.Tag, logData.Level, logData.Message)

		c.JSON(http.StatusOK, gin.H{"message": "Log recorded"})
	})

	// اجرای سرور
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

