package middlewares

import (
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/vars"

	"github.com/gin-gonic/gin"
)

func AllowRoles(log logger.LoggerHandler, allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesString := c.MustGet("roles")
		userRoles := make([]string, len(rolesString.([]interface{})))
		for i, v := range rolesString.([]interface{}) {
			userRoles[i] = v.(string)
		}

		if !includes(userRoles, allowedRoles) {
			log.Error("Not allowed to access")
			c.JSON(vars.STATUS_AUTH_ERROR_CODE, gin.H{
				"message": vars.STATUS_AUTH_ERROR_STRING,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// includes function
func includes(userRoles []string, allowedRoles []string) bool {
	for _, allowedRole := range allowedRoles {
		for _, userRole := range userRoles {
			if allowedRole == userRole {
				return true
			}
		}
	}
	return false
}
