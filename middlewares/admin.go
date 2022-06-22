package middlewares

import (
	"go-api-starter-kit/helpers"
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/vars"

	"github.com/gin-gonic/gin"
)

func AllowAdmin(log logger.LoggerHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		globaladmin, err := helpers.GetClaimsBool(c.MustGet("globaladmin"))
		if err != nil {
			c.JSON(vars.STATUS_AUTH_ERROR_CODE, gin.H{
				"message": vars.STATUS_AUTH_ERROR_STRING,
			})
			c.Abort()
			return
		}

		if !globaladmin {
			c.JSON(vars.STATUS_AUTH_ERROR_CODE, gin.H{
				"message": vars.STATUS_AUTH_ERROR_STRING,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
