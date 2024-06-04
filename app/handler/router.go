package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h handler) RegisterHandlers(router *mux.Router) *mux.Router {
	// Define a HTTP endpoint called GET /range-fizzbuzz
	router.HandleFunc("/range-fizzbuzz", func(w http.ResponseWriter, r *http.Request) {
		goroutineSem := make(chan struct{}, 1000) // maximum goroutine
		defer close(goroutineSem)

		select {
		case goroutineSem <- struct{}{}:
			h.handleSingleFizzBuzzWithRange(w, r)
			<-goroutineSem
		default:
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Server too busy, please try again"))
		}
	}).Methods(http.MethodGet)

	return router
}
