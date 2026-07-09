package utils

import (
	"crypto/rand"
	"math/big"
	"time"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62(num int64) string {
	if num <= 0 {
		return "0"
	}
	b := make([]byte, 0, 11) // int64 최대값도 11자리면 충분
	for num > 0 {
		b = append(b, base62Chars[num%62])
		num /= 62
	}
	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func Uid() string {
	rn, err := rand.Int(rand.Reader, big.NewInt(52))
	if err != nil {
		panic(err) // crypto/rand 실패는 복구 불가
	}
	prefix := Base62(rn.Int64() + 10) // 'a'~'Z' 한 글자, 숫자 시작 방지
	micro := Base62(time.Now().UnixMicro())

	// 같은 마이크로초 내 충돌 방지용 랜덤 suffix (62^3 ≈ 238,000 조합)
	//	sn, err := rand.Int(rand.Reader, big.NewInt(62*62*62))
	//	if err != nil {
	//		panic(err)
	//	}
	//	suffix := Base62(sn.Int64())

	return prefix + micro
}
