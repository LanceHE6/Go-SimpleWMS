package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// CreateVerifyCode 生成6位验证码
func CreateVerifyCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
