package router

import (
	"net/http"
	"win/controler/cmd/api"
	"win/controler/cmd/middleware"
)

func Router(api *api.ApiConfig) http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.Hello)
	mux.HandleFunc("/api/tx", api.TxHandler)
	mux.HandleFunc("/api/env", api.EnvoiceAllHandler)
	mux.HandleFunc("/api/env/num", api.EnvoiceOneHandler)

	return middleware.Cors(mux)

}
