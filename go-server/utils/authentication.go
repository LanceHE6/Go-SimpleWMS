package utils

import (
	"Go_simpleWMS/config"
	"Go_simpleWMS/database/model"
	db2 "Go_simpleWMS/database/my_db"
	"Go_simpleWMS/utils/response"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(config.ServerConfig.SERVER.SECRET_KEY) // 用于签名的密钥

// MyClaims 自定义载荷内容
type MyClaims struct {
	jwt.StandardClaims
	Uid        string `json:"uid"`
	Permission int    `json:"permission"`
	CreatedAT  string `json:"created_at"`
	SessionID  string `json:"session_id"`
}

// GenerateToken 生成一个token
func GenerateToken(id string, permission int, createdAt string, sessionID string) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		Uid:        id,
		Permission: permission,
		CreatedAT:  createdAt,
		SessionID:  sessionID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthMiddleware 基础鉴权的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Response(101, "No Authorization header provided", nil))
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, response.Response(102, "Invalid token", nil))
			c.Abort()
			return
		}
		myClaims, err := GetUserInfoByContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse(103, "Invalid token", err.Error()))
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(bearerToken[1], &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			var ve *jwt.ValidationError
			if errors.As(err, &ve) {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					// Token is expired
					c.JSON(http.StatusUnauthorized, response.Response(104, "Token is expired", nil))
				} else {
					// Other errors
					c.JSON(http.StatusUnauthorized, response.Response(105, "Invalid token", nil))
				}
			}
			c.Abort()
			return
		}

		// 判断是否在数据库中
		var user model.User
		db := db2.GetMyDbConnection()
		err = db.Where("uid=? and session_id=?", myClaims.Uid, myClaims.SessionID).First(&user).Error

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
				"code":    106,
			})
			c.Abort()
			return
		}

		if _, ok := token.Claims.(*MyClaims); ok && token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
				"code":    107,
			})
			c.Abort()
			return
		}
	}
}

// IsAdminMiddleware 管理员鉴权中间件 >=2
func IsAdminMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 根据token判断permission是否大于等于2
		myClaims, err := GetUserInfoByContext(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, response.Response(108, "Invalid token", nil))
			context.Abort()
			return
		}

		if myClaims.Permission >= 2 {
			context.Next()
		} else {
			context.JSON(http.StatusForbidden, response.Response(110, "Permission denied", nil))
			context.Abort()
			return
		}
	}
}

// IsSuperAdminMiddleware 超级管理员鉴权的中间件 ==3
func IsSuperAdminMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 根据token判断permission是否为3
		myClaims, err := GetUserInfoByContext(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, response.Response(109, "Invalid token", nil))
			context.Abort()
			return
		}

		if myClaims.Permission == 3 {
			context.Next()
		} else {
			context.JSON(http.StatusForbidden, response.Response(111, "Permission denied", nil))
			context.Abort()
			return
		}
	}
}

// GetUserInfoByContext 通过context获取载荷信息
func GetUserInfoByContext(context *gin.Context) (MyClaims, error) {
	authHeader := context.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	// 解析token
	claims := MyClaims{}
	_, err := jwt.ParseWithClaims(bearerToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	// 从token中获取载荷数据
	return claims, err
}
