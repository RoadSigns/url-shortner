package app

import (
	"github.com/roadsigns/url-shortner/controllers/polo"
	"github.com/roadsigns/url-shortner/controllers/shorten"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/shorten", shorten.Execute)
}
