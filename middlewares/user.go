package middlewares

import (
	"database/sql"
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/vars"

	"github.com/gin-gonic/gin"
)

type usermodel struct {
	uid    sql.NullInt64
	aid    sql.NullInt64
	locked sql.NullBool
}

func User(db *sql.DB, log logger.LoggerHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var user usermodel

		email := c.MustGet("email")

		err = db.QueryRow("SELECT id, account_id, locked FROM users WHERE email=$1;", email).Scan(&user.uid, &user.aid, &user.locked)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Error(err.Error())
				c.JSON(vars.STATUS_NOTFOUND_USER_ERROR_CODE, gin.H{
					"message": vars.STATUS_NOTFOUND_USER_ERROR_STRING,
				})
				c.Abort()
				return
			}
			log.Error(err.Error())
			c.JSON(vars.STATUS_DB_ERROR_CODE, gin.H{
				"message": vars.STATUS_DB_ERROR_STRING,
			})
			c.Abort()
			return
		}

		if user.locked.Valid && user.locked.Bool {
			c.JSON(vars.STATUS_AUTH_LOCKED_ERROR_CODE, gin.H{
				"message": vars.STATUS_AUTH_LOCKED_ERROR_STRING,
			})
			c.Abort()
			return
		}

		if user.uid.Valid {
			c.Set("uid", user.uid.Int64)
		} else {
			log.Error(err.Error())
			c.JSON(vars.STATUS_NOTFOUND_ERROR_CODE, gin.H{
				"message": vars.STATUS_NOTFOUND_ERROR_STRING,
			})
			c.Abort()
			return
		}

		if user.aid.Valid {
			c.Set("aid", user.aid.Int64)
		} else {
			log.Error(err.Error())
			c.JSON(vars.STATUS_NOTFOUND_ERROR_CODE, gin.H{
				"message": vars.STATUS_NOTFOUND_ERROR_STRING,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
