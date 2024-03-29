package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(SecretKey) // 用于签名的密钥

// 自定义载荷内容
type myClaims struct {
	jwt.StandardClaims
	Uid        string `json:"uid"`
	Permission int    `json:"permission"`
}

// GenerateToken 生成一个token
func GenerateToken(id string, permission int) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		Uid:        id,
		Permission: permission,
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
		uid, err := GetUidByContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(bearerToken[1], &myClaims{}, func(token *jwt.Token) (interface{}, error) {
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

		// 判断uid是否在数据库中
		tx, _ := GetDbConnection()
		var isExist int
		err = tx.QueryRow("SELECT count(*) from user where uid=?", uid).Scan(&isExist)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot get the number of uid for this uid",
				"detail": err.Error(),
			})
			c.Abort()
			return
		}
		if isExist <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		if _, ok := token.Claims.(*myClaims); ok && token.Valid {

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
		uid, err := GetUidByContext(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			context.Abort()
			return
		}

		tx, err := GetDbConnection()

		if tx == nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Cannot begin transaction",
				"detail": err.Error(),
			})
			context.Abort()
			return
		}

		var permission string
		err = tx.QueryRow("SELECT permission FROM user WHERE uid=?", uid).Scan(&permission)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
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

// GetUidByContext 获取用户id
func GetUidByContext(context *gin.Context) (string, error) {
	authHeader := context.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	// 解析token
	claims := &myClaims{}
	_, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	// 从token中获取载荷数据
	uid := claims.Uid
	return uid, err
}
