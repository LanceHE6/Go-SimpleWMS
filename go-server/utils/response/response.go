package response

import "github.com/gin-gonic/gin"

// Response 构造返回信息
func Response(code int, message string, data gin.H) gin.H {
	return gin.H{
		"msg":  message,
		"data": data,
		"code": code,
	}
}

// ErrorResponse 构造错误返回信息
func ErrorResponse(code int, message string, errStr string) gin.H {
	return Response(code, message, gin.H{"err": errStr})
}

var IncorrectParamsStr = "Missing parameters or incorrect format"

func MissingParamsResponse(err error) gin.H {
	return ErrorResponse(401, IncorrectParamsStr, err.Error())
}
