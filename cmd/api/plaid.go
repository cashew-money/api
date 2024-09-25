package main

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/plaid/plaid-go/plaid"
)

func SandboxPublicTokenCreate(app *Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		sandboxPublicTokenResp, _, err := app.PlaidClient.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
			*plaid.NewSandboxPublicTokenCreateRequest(
				"ins_109508",
				[]plaid.Products{plaid.PRODUCTS_TRANSACTIONS},
			),
		).Execute()
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		publicToken := sandboxPublicTokenResp.GetPublicToken()

		data := envelope{
			"public_token": publicToken,
		}

		err = app.writeJSON(w, http.StatusCreated, data, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
}
