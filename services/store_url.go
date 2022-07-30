package services

import "github.com/roadsigns/url-shortner/shortener"

type StoreUrl struct {
	Shortener shortener.Shortener
}

func (s StoreUrl) Save(longUrl string) (string, error) {
	// Encode the long url to a shorter version
	shortUrl := s.Shortener.Shorten(longUrl)
	// Pass the URL to the repository

	return shortUrl, nil
}
