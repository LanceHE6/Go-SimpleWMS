package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(SecretKey) // 用于签名的密钥

// GenerateToken 生成一个token
func GenerateToken(id string) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   id,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthMiddleware 是一个用于鉴权的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No Authorization header provided"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(bearerToken[1], &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			var ve *jwt.ValidationError
			if errors.As(err, &ve) {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					// Token is expired
					c.JSON(http.StatusUnauthorized, gin.H{"message": "Expired token"})
				} else {
					// Other errors
					c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
				}
			}
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			c.Set("userID", claims.Subject)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

	}
}

func IsSuperAdminMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 根据token判断permission是否为3
		token := context.GetHeader("Authorization")
		bearerToken := strings.Split(token, " ")[1]
		fmt.Println(bearerToken)
		db := GetDbConnection()

		// 开始一个新的事务
		tx, err := db.Begin()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot begin authorization transaction"})
			context.Abort()
			return
		}
		defer func(tx *sql.Tx) {
			err := tx.Rollback()
			if err != nil {

			}
		}(tx) // 如果出错，回滚事务

		var permission string
		err = tx.QueryRow("SELECT permission FROM user WHERE token=?", bearerToken).Scan(&permission)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get target user permission"})
			context.Abort()
			return
		}
		if permission == "3" {
			context.Next()

		} else {
			context.JSON(http.StatusForbidden, gin.H{"message": "Permission denied"})
			context.Abort()
			return
		}
	}
}
