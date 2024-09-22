package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Healthcheck(app *Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "status: available")
	}
}
