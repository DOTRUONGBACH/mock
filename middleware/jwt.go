package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy thông tin về token từ header của yêu cầu
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Xác thực token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra loại algorithm của token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			// Trả về key bí mật
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Lưu thông tin về token vào context của yêu cầu
		c.Set("token", token)

		// Gọi tiếp theo trong chuỗi middleware
		c.Next()
	}
}
