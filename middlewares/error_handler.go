package middlewares

import (
	"gin-blueprint/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler - Global error handling middleware
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Handler'dan dönen hataları kontrol et
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// AppError ise
			if appErr, ok := err.Err.(*utils.AppError); ok {
				if appErr.Details != nil {
					utils.ErrorResponseWithDetails(
						c,
						appErr.StatusCode,
						appErr.Code,
						appErr.Message,
						appErr.Details,
					)
				} else {
					utils.ErrorResponse(
						c,
						appErr.StatusCode,
						appErr.Code,
						appErr.Message,
					)
				}
				return
			}

			// Genel hata
			log.Printf("Unhandled error: %v", err.Err)
			utils.ErrorResponse(
				c,
				http.StatusInternalServerError,
				"INTERNAL_ERROR",
				"An unexpected error occurred",
			)
		}
	}
}
