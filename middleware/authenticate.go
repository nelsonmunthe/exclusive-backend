package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		tokenString := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)

		authData, valid := validateJwt(c, tokenString)

		if valid {
			c.Set("authData", authData)
			c.Next()
			return
		}

		response := defaultUnauthorizedResponse()
		response.ResponseTime = fmt.Sprint(time.Since(start).Milliseconds(), " ms.")
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
	}
}

func GetAuthData(c *gin.Context) map[string]interface{} {
	authDataValue, exists := c.Get("authData")
	if !exists || authDataValue == nil {
		return nil
	}

	return authDataValue.(map[string]interface{})
}

func validateJwt(c *gin.Context, tokenString string) (map[string]interface{}, bool) {

	var secretKey = []byte("secret-key")
	token, err := parseJwt(tokenString, secretKey)
	if err != nil {
		log.Println("middleware.parseJwt:", err)
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	valid := ok && token.Valid
	var authData map[string]interface{}

	if valid {
		authData = claims["data"].(map[string]interface{})
	}

	return authData, valid
}

func parseJwt(tokenString string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		fmt.Println("secret", secret)
		return secret, nil
	})
}

func defaultUnauthorizedResponse() unauthorizedResponse {
	return unauthorizedResponse{
		Success:      false,
		Message:      "Auth Failed",
		ResponseTime: "",
	}
}

func GetAuthDataStruct(c *gin.Context) (AuthData, error) {
	authData := AuthData{}
	err := authData.LoadFromMap(GetAuthData(c))
	return authData, err
}

type unauthorizedResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	MessageTitle *any   `json:"messageTitle"`
	Data         *any   `json:"data"`
	ResponseTime string `json:"responseTime"`
}
