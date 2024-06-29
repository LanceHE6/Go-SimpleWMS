package utils

import "github.com/google/uuid"

// GenerateUuid 生成指定长度的UUID
func GenerateUuid(length int) string {
	newUuid := uuid.New().String()
	return newUuid[len(newUuid)-length:]
}
