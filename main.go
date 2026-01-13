package main

import (
	"gin-blueprint/database"
	"gin-blueprint/middlewares"
	"gin-blueprint/validators"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type HelleResponse struct {
	Message any `json:"message"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env dosyası bulunamadı, environment variables kullanılacak")
	}
	database.Connect()
	database.Migrate()

	r := gin.New()
	r.Use(middlewares.CustomLogger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORS())
	r.Use(middlewares.ErrorHandler())

	rateLimiter := middlewares.NewRateLimiter(100, time.Minute)
	r.Use(rateLimiter.Middleware())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("username", validators.ValidateUsername)
		v.RegisterValidation("strongpassword", validators.ValidateStrongPassword)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})
	setupRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, HelleResponse{Message: 1})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
