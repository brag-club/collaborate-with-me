package middleware

import (
	appContext "github.com/BRAVO68WEB/collaborate-with-me/packages/api/utils"

	"github.com/gin-gonic/gin"
)

// GinContextToContext is the middleware to convert gin context to go context
func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appContext.CreateContextFromGinContext(c)
		ctx = appContext.WithRequest(ctx, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
