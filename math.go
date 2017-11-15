// Package utils go实用库
// RSA的签名及验签
// SHA256 With RSA
package utils

import "math"

// Round 四舍五入
func Round(v float64) float64 {
	i, f := math.Modf(v)
	if f >= 0.5 {
		return i + 1.0
	}
	return i
}

// Floor 舍弃法
func Floor(v float64) float64 {
	i, f := math.Modf(v)
	if f >= 0.1 {
		return i - f
	}
	return i
}

// Ceil 进一法
func Ceil(v float64) float64 {
	i, f := math.Modf(v)
	if f >= 0.1 {
		return i + 1.0
	}
	return i
}
