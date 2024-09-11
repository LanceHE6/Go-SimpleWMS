package utils

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

// HashPsw 将密码进行哈希
func HashPsw(psw string) string {
	var hashInstance hash.Hash // 定义哈希实例

	hashInstance = sha256.New()

	hashInstance.Write([]byte(psw)) // 将字符串转换为字节数组，写入哈希对象

	bytes := hashInstance.Sum(nil)  // 哈希值追加到参数后面，只获取原始值，不用追加，用nil，返回哈希值字节数组
	return fmt.Sprintf("%x", bytes) // 格式化输出哈希值
}

// CheckPsw 检查密码是否正确
func CheckPsw(hashedPsw, psw string) bool {
	return hashedPsw == HashPsw(psw)
}
