package middleware

import (
	"fmt"
	"net/http"
	"strings"
	resultdto "test-absensi/dto/result"
	jwtToken "test-absensi/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "unauthorized 1"})
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "unathorized 2"})
			return
		}

		c.Set("userLogin", claims)
		fmt.Println("userLogin")
		next(c) // agar bisa di teruskan ke handler selanjutnya
	}
}