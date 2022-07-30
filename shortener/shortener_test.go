package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase62Shortener_Shorten(t *testing.T) {
	shortener := Base62Shortener{}

	link1 := shortener.Shorten("https://google.co.uk")
	assert.Equal(t, link1, "rVnLvNmL")

	link2 := shortener.Shorten("https://amazon.com")
	assert.Equal(t, link2, "t92Yu42b")

	link3 := shortener.Shorten("https://go.dev/")
	assert.Equal(t, link3, "vYXZk5yb")
}
