package middlewares

import (
	"crypto/rsa"
	"errors"
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/vars"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(keysPath string, keyPrefix string, log logger.LoggerHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getAuthToken(c, log)
		if err != nil {
			log.Error("No token error")
			c.JSON(err.(*vars.StatusErr).Code(), gin.H{
				"message": err.(*vars.StatusErr).Error(),
			})
			return
		}

		var parsedToken *jwt.Token
		pem, err := getPem(log, keysPath, keyPrefix)
		if err != nil {
			log.Error("No token error")
			c.JSON(err.(*vars.StatusErr).Code(), gin.H{
				"message": err.(*vars.StatusErr).Error(),
			})
			return
		}

		parsedToken, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return pem, nil })
		if err != nil {
			log.Error("No token error")
			c.JSON(vars.STATUS_SERVER_ERROR_CODE, gin.H{
				"message": vars.STATUS_SERVER_ERROR_STRING,
			})
			c.Abort()
			return
		}

		err = getClaims(c, parsedToken, log)
		if err != nil {
			log.Error("No token error")
			c.JSON(err.(*vars.StatusErr).Code(), gin.H{
				"message": err.(*vars.StatusErr).Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func getClaims(c *gin.Context, token *jwt.Token, log logger.LoggerHandler) error {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		c.Set("email", claims["email"])
		// TODO migrate roles to user.go middleware
		c.Set("roles", claims["roles"])
	} else {
		return vars.StatusServerError(errors.New("getClaimsError"))
	}
	return nil
}

func getAuthToken(c *gin.Context, log logger.LoggerHandler) (string, error) {
	var token string

	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 {
		log.Error("Header parse error")
		return token, vars.StatusServerError(errors.New("headerIssue"))
	}

	token = strings.TrimSpace(splitToken[1])

	return token, nil
}

func getPem(log logger.LoggerHandler, keysPath string, keyPrefix string) (*rsa.PublicKey, error) {
	var pubKeyPath = keysPath + "/" + keyPrefix + "_jwtRS256.key.pub"

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Error(err.Error())
		return nil, vars.StatusServerError(err)
	}

	verifiedPem, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Error(err.Error())
		return nil, vars.StatusServerError(err)
	}

	return verifiedPem, nil
}
