package main

import (
	"url-stortener/internal/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		panic(err)
	}

	if err := application.Run(); err != nil {
		panic(err)
	}
}
