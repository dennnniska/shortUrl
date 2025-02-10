package random

import (
	"math/rand/v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	capital, lowercase, underlining, numb := false, false, false, false
	for !capital || !lowercase || !underlining || !numb {
		capital, lowercase, underlining, numb = false, false, false, false
		for i, cache, remain := n-1, rand.Int64(), letterIdxMax; i >= 0; {
			if remain == 0 {
				cache, remain = rand.Int64(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
				b[i] = letterBytes[idx]
				i--
			}
			cache >>= letterIdxBits
			remain--
		}
		for i := 0; i < len(b); i++ {
			if b[i] >= 'a' && b[i] <= 'z' {
				lowercase = true
			}
			if b[i] >= '0' && b[i] <= '9' {
				numb = true
			}
			if b[i] >= 'A' && b[i] <= 'Z' {
				capital = true
			}
			if b[i] == '_' && underlining {
				underlining = false
				break
			}
			if b[i] == '_' && !underlining {
				underlining = true
			}
		}
	}
	return string(b)
}
