package handlers

import (
	"net/http"

	"github.com/cashew-money/api/cmd/api/config"
	"github.com/julienschmidt/httprouter"
)

func Healthcheck(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		res := envelope{
			"status": "available",
		}

		err := writeJSON(w, http.StatusOK, res, nil)
		if err != nil {
			ServerErrorResponse(env, w, r, err)
		}
	}
}
