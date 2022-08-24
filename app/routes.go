package app

import (
	"github.com/roadsigns/url-shortner/controllers/polo"
	"github.com/roadsigns/url-shortner/controllers/redirect"
	"github.com/roadsigns/url-shortner/controllers/shorten"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.GET("/:shortUrl", redirect.Execute)
	router.POST("/shorten", shorten.Execute)
}
