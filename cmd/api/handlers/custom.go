package handlers

import (
	"net/http"

	"github.com/cashew-money/api/cmd/api/config"
)

func NotFound(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NotFoundResponse(env, w, r)
	}
}

func MethodNotAllowed(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		MethodNotAllowedResponse(env, w, r)
	}
}
