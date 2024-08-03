package utils

import "github.com/google/uuid"

// GenerateUUID 生成指定长度的UUID 长度为-1时，生成默认长度32的UUID
func GenerateUUID(length int) string {
	// length为-1时，生成默认长度32的UUID
	newUUID := uuid.New().String()
	if length == -1 {
		return newUUID
	}
	return newUUID[len(newUUID)-length:]
}
