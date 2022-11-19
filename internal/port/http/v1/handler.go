package v1

import (
	"github.com/go-chi/chi/v5"
	"grain-acceptance/internal/app"
	"net/http"
)

type handler struct {
	app app.Application
}

func NewHandler(app app.Application, router chi.Router) http.Handler {
	return HandlerFromMux(handler{
		app: app,
	}, router)
}
