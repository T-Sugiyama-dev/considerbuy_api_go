package config

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORS() gin.HandlerFunc {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	clientOrigin := os.Getenv("CLIENT_ORIGIN")

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", clientOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
