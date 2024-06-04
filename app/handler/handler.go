package handler

import (
	"github.com/gorilla/mux"
	"github.com/kevinsudut/go-learn/app/usecase"
)

type handler struct {
	usecase usecase.UsecaseItf
}

type HandlerItf interface {
	RegisterHandlers(router *mux.Router) *mux.Router
}

func Init() HandlerItf {
	return &handler{
		usecase: usecase.Init(),
	}
}
