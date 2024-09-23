package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes(app *Application) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.GET("/v1/healthcheck", Healthcheck(app))

	return app.recoverPanic(router)
}
