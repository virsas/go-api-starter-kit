package test

import (
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt"
)

// Claims struct
type Claims struct {
	UserEmail string   `json:"email"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}

func SetupJWT(keysPath string, email string, roles []string) string {
	var privKeyPath = keysPath + "/test_jwtRS256.key"

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		panic(err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	claims := Claims{
		email,
		roles,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, err := token.SignedString(signKey)
	if err != nil {
		panic(err)
	}

	return signedString
}
