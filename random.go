package utils

import (
	"math/rand"
	"time"
)

// Random 随机生成指定范围的数
func Random(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min) + min
}

// RandomNumber 随机生成一个数
func RandomNumber() int64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
}

// RandomStr 随机指定长度的字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	byteStr := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, byteStr[r.Intn(len(byteStr))])
	}
	return string(result)
}
