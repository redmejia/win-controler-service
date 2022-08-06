package router

import (
	"net/http"
	"win/controler/cmd/api"
)

func Router(api *api.ApiConfig) http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.Hello)
	mux.HandleFunc("/api/tx", api.TxControler)

	return mux

}
