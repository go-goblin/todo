package dependencies

import (
	"url-stortener/internal/service"
)

type Dependencies struct {
	UrlShortenerService service.UrlShortenerService
}

func New(urlShortenerService service.UrlShortenerService) *Dependencies {
	return &Dependencies{
		UrlShortenerService: urlShortenerService,
	}
}
