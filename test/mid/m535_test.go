package test

import (
	"fmt"
	"math/rand"
	"testing"
)

type Codec struct {
	origin2Tiny map[string]string
	tiny2Origin map[string]string
}

func Constructor535() Codec {
	return Codec{
		origin2Tiny: map[string]string{},
		tiny2Origin: map[string]string{},
	}
}

const baseStr = "qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
const baseLen = 26*2 + 10
const k = 6

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	if v, ok := c.origin2Tiny[longUrl]; ok {
		return v
	}
	cs := make([]byte, k)
	for i := 0; i < k; i++ {
		cs[i] = baseStr[rand.Intn(baseLen)]
	}
	str := string(cs)
	c.origin2Tiny[longUrl] = str
	c.tiny2Origin[str] = longUrl
	return str
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.tiny2Origin[shortUrl]
}

func Test535(t *testing.T) {
	code := Constructor535()
	tiny := code.encode("abc")
	fmt.Println(tiny)
	fmt.Println(code.decode(tiny))
}
