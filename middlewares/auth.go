package middlewares

import (
	"crypto/rsa"
	"errors"
	"go-api-starter-kit/config"
	"go-api-starter-kit/utils/logger"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(validateJWT bool, log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getAuthToken(c, log)
		if err != nil {
			c.JSON(err.(*config.CustErr).Code(), gin.H{
				"message": err.(*config.CustErr).Error(),
			})
			return
		}

		var parsedToken *jwt.Token
		pem, err := getPem(log)
		if err != nil {
			c.JSON(err.(*config.CustErr).Code(), gin.H{
				"message": err.(*config.CustErr).Error(),
			})
			return
		}

		parsedToken, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return pem, nil })
		if err != nil {
			log.Error("No token error")
			c.JSON(config.SERVER_ERROR, gin.H{
				"message": config.SERVER_STRING,
			})
			c.Abort()
			return
		}

		err = getClaims(c, parsedToken, log)
		if err != nil {
			c.JSON(err.(*config.CustErr).Code(), gin.H{
				"message": err.(*config.CustErr).Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func getClaims(c *gin.Context, token *jwt.Token, log *logger.Logger) error {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		c.Set("uid", claims["uid"])
	} else {
		return config.NewServerError(errors.New("getClaimsError"))
	}
	return nil
}

func getAuthToken(c *gin.Context, log *logger.Logger) (string, error) {
	var token string

	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 {
		log.Error("Header parse error")
		return token, config.NewServerError(errors.New("headerIssue"))
	}

	token = strings.TrimSpace(splitToken[1])

	return token, nil
}

func getPem(log *logger.Logger) (*rsa.PublicKey, error) {
	currentDir, _ := os.Getwd()
	var pubKeyPath = currentDir + "/keys/jwtRS256.key.pub"

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Error(err.Error())
		return nil, config.NewServerError(err)
	}

	verifiedPem, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Error(err.Error())
		return nil, config.NewServerError(err)
	}

	return verifiedPem, nil
}
