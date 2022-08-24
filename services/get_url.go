package services

import "github.com/roadsigns/url-shortner/repositories"

type GetUrl struct {
	Repository repositories.UrlRepository
}

func (s GetUrl) Get(shortUrl string) (string, error) {
	urlData, err := s.Repository.Get(shortUrl)
	if err != nil {
		return "", err
	}

	return urlData.LongUrl, nil
}
