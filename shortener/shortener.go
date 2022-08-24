package shortener

import (
	"github.com/jxskiss/base62"
)

type Shortener interface {
	Shorten(url string) string
}

type Base62Shortener struct {
}

func (s Base62Shortener) Shorten(url string) string {
	return string(base62.Encode([]byte(url)))[:8]
}
