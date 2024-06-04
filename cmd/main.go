package main

import (
	"errors"
	"net/http"

	"github.com/kevinsudut/go-learn/app"
	"github.com/kevinsudut/go-learn/pkg/lib/log"
)

func main() {
	log.Init()

	err := app.Init()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("app.Init", err)
	}
}
