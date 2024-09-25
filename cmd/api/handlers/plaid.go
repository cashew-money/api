package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/cashew-money/api/cmd/api/config"
	"github.com/julienschmidt/httprouter"
	"github.com/plaid/plaid-go/plaid"
)

func SandboxPublicTokenCreate(env *config.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		sandboxPublicTokenResp, _, err := env.PlaidClient.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
			*plaid.NewSandboxPublicTokenCreateRequest(
				"ins_109508",
				[]plaid.Products{plaid.PRODUCTS_TRANSACTIONS},
			),
		).Execute()
		if err != nil {
			ServerErrorResponse(env, w, r, err)
			return
		}

		publicToken := sandboxPublicTokenResp.GetPublicToken()

		data := envelope{
			"public_token": publicToken,
		}

		err = writeJSON(w, http.StatusCreated, data, nil)
		if err != nil {
			ServerErrorResponse(env, w, r, err)
		}
	}
}
