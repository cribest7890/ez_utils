package random

import (
	"crypto/rand"
	"math/big"
)

func Int(min, max int) int {
	if min >= max {
		return min
	}
	bg := big.NewInt(int64(max - min + 1))
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		return min
	}
	return min + int(n.Int64())
}

func Float64() float64 {
	bg := big.NewInt(1 << 53)
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		return 0.0
	}
	return float64(n.Int64()) / (1 << 53)
}

func Float64Range(min, max float64) float64 {
	if min >= max {
		return min
	}
	return min + Float64()*(max-min)
}

func Bool() bool {
	return Int(0, 1) == 1
}

func String(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[Int(0, len(charset)-1)]
	}
	return string(b)
}

func Choice[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[Int(0, len(slice)-1)]
}

func Shuffle[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	for i := len(result) - 1; i > 0; i-- {
		j := Int(0, i)
		result[i], result[j] = result[j], result[result[i]]
	}
	return result
}
