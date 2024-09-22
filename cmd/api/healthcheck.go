package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Healthcheck(app *Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data := map[string]string{
			"status": "available",
		}

		err := app.writeJSON(w, http.StatusOK, data, nil)
		if err != nil {
			app.Logger.Error(err.Error())
			http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		}
	}
}
