package app

import (
	"github.com/roadsigns/url-shortner/internal/app/url-shortner/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
}
