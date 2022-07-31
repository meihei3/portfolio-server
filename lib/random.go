package lib

import (
	"math/rand"
)

const (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lettersBit = 6
	chunkMask  = 1<<lettersBit - 1
	chunk      = 63 / lettersBit
)

func GenerateRandomStr(n int) string {
	b := make([]byte, n)
	r, c := rand.Int63(), chunk
	for i := n - 1; i >= 0; {
		if c == 0 {
			r, c = rand.Int63(), chunk
		}
		idx := int(r & chunkMask)
		if idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		r >>= lettersBit
		c--
	}
	return string(b)
}
