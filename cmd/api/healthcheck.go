package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Healthcheck(app *Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		env := envelope{
			"status": "available",
		}

		err := app.writeJSON(w, http.StatusOK, env, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
}
