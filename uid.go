package utils

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"
)

func Base62(num int64) string {
	CharacterSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 0)

	// loop as long the num is bigger than zero
	for num > 0 {
		r := math.Mod(float64(num), float64(62))
		num /= 62
		b = append([]byte{CharacterSet[int(r)]}, b...)
	}

	return string(b)
}

func Uid() string {
	now := time.Now()
	rn, _ := rand.Int(rand.Reader, big.NewInt(52))
	prefix := Base62(rn.Int64() + 10)
	sec := now.UnixNano() / 1000
	uid := Base62(sec)

	return prefix + uid
}
