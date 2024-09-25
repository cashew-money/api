package handlers

import (
	"fmt"
	"net/http"

	"github.com/cashew-money/api/cmd/api/config"
)

func logError(env *config.Env, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	env.Logger.Error(err.Error(), "method", method, "uri", uri)
}

func ErrorResponse(env *config.Env, w http.ResponseWriter, r *http.Request, status int, message any) {
	res := envelope{"error": message}

	err := writeJSON(w, status, res, nil)
	if err != nil {
		logError(env, r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(env *config.Env, w http.ResponseWriter, r *http.Request, err error) {
	logError(env, r, err)

	message := "the server encountered a problem and could not process your request"
	ErrorResponse(env, w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(env *config.Env, w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	ErrorResponse(env, w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(env *config.Env, w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	ErrorResponse(env, w, r, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(env *config.Env, w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(env, w, r, http.StatusBadRequest, err.Error())
}
