package middlewares

import (
	"github.com/artur244/first-go-rest-api/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const Bearer_schema = "Bearer "

		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJwtService().ValidadeToken(token) {
			ctx.AbortWithStatus(401)
		}
	}
}
