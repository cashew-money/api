package main

import (
	"net/http"

	"github.com/cashew-money/api/cmd/api/config"
	"github.com/cashew-money/api/cmd/api/handlers"
	"github.com/julienschmidt/httprouter"
)

func routes(env *config.Env) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(handlers.NotFound(env))
	router.MethodNotAllowed = http.HandlerFunc(handlers.MethodNotAllowed(env))

	router.GET("/v1/healthcheck", handlers.Healthcheck(env))

	router.POST("/v1/plaid/sandbox/public_token/create", handlers.SandboxPublicTokenCreate(env))

	return recoverPanic(env, router)
}
