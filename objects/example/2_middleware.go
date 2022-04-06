package example

import (
	"bytes"
	"encoding/json"
	"go-api-starter-kit/helpers"
	"go-api-starter-kit/utils/config"
	"go-api-starter-kit/utils/logger"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func validateExample(log logger.LoggerHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		data, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		var example ExampleInput
		err = json.Unmarshal(data, &example)
		if err != nil {
			c.JSON(config.REQUEST_ERROR, gin.H{
				"message": config.REQUEST_STRING,
			})
			c.Abort()
			return
		}

		validate := validator.New()
		validate.RegisterValidation("alphanumspace", helpers.AlphaNumSpaceValid)
		validate.RegisterValidation("alphaspace", helpers.AlphaSpaceValid)
		err = validate.Struct(&example)
		if err != nil {
			log.Error(err.Error())
			c.JSON(config.VALID_ERROR, gin.H{
				"message": config.VALID_STRING,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
