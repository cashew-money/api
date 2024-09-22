package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes(app *Application) http.Handler {
	router := httprouter.New()

	router.GET("/v1/healthcheck", Healthcheck(app))

	return router
}
