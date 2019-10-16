package handlers

import (
	"math"
	"math/rand"
	"time"
)

const (
	base62   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = 62.0
	idLength = 7.0
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateBase62() string {
	n := rand.Intn(int(math.Pow(base, idLength)))

	// https://github.com/kare/base62/blob/master/base62.go
	b := make([]byte, 0)
	for n > 0 {
		r := math.Mod(float64(n), float64(base))
		n /= base
		b = append([]byte{base62[int(r)]}, b...)
	}
	return string(b)
}
