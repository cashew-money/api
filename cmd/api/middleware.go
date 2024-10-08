package main

import (
	"fmt"
	"net/http"

	"github.com/cashew-money/api/cmd/api/config"
	"github.com/cashew-money/api/cmd/api/handlers"
)

func recoverPanic(env *config.Env, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				handlers.ServerErrorResponse(env, w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
