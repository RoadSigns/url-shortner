package services

import (
	"github.com/roadsigns/url-shortner/repositories"
	"github.com/roadsigns/url-shortner/shortener"
)

type StoreUrl struct {
	Shortener  shortener.Shortener
	Repository repositories.UrlRepository
}

func (s StoreUrl) Save(longUrl string) (string, error) {
	shortUrl := s.Shortener.Shorten(longUrl)

	err := s.Repository.Save(repositories.UrlData{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	})

	if err != nil {
		return "", err
	}

	return shortUrl, nil
}
