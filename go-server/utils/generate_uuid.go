package utils

import "github.com/google/uuid"

func GenerateUuid(length int) string {
	newUuid := uuid.New().String()
	return newUuid[len(newUuid)-length:]
}
